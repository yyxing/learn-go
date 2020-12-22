package pool

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

var (
	defaultPoolCapacity int32 = 10240
)

const (
	DefaultCleanIntervalTime = time.Duration(6000) * time.Millisecond
	// OPENED represents that the pool is opened.
	OPENED = iota

	// CLOSED represents that the pool is closed.
	CLOSED
)

type Pool struct {
	// 连接池大小
	capacity int32
	// 是否可阻塞
	isBlock bool
	// 正在运行数量
	running int32
	// 工作协程 也就是工作的channel
	workers workerArray
	// 超时时间
	expiryDuration time.Duration
	// 协程池状态
	state int32
	// 同步操作锁
	lock sync.Locker
	// 状态通知器
	cond *sync.Cond
	// go自带的缓存池 GC时清空local
	workerCache *sync.Pool
	// 阻塞的数量
	blockingNum int32
}

func Default() *Pool {
	pool := &Pool{
		capacity: defaultPoolCapacity,
		isBlock:  false,
	}
	return pool
}

func New(capacity int32, isBlock bool, expire int64) (*Pool, error) {
	var expireDuration time.Duration
	if expire <= 0 {
		expireDuration = DefaultCleanIntervalTime
	} else {
		expireDuration = time.Duration(expire) * time.Millisecond
	}
	pool := &Pool{
		capacity:       capacity,
		isBlock:        isBlock,
		expiryDuration: expireDuration,
		lock:           NewSpinLock(),
	}
	pool.workerCache.New = func() interface{} {
		return &worker{
			pool: pool,
			task: make(chan func(), 1),
		}
	}
	pool.cond = sync.NewCond(pool.lock)
	go pool.purgePeriodically()
	return pool, nil
}

func (p *Pool) Submit(task func()) error {
	if atomic.LoadInt32(&p.state) == CLOSED {
		return errors.New("pool is closed")
	}
	var w *worker
	if w = p.retrieveWorker(); w == nil {
		return errors.New("无可用worker")
	}
	w.task <- task
	return nil
}

// 返回一个可用的worker 运行task
func (p *Pool) retrieveWorker() (w *worker) {
	// 从系统Pool中取一个worker并且运行 若不开协程池的大小 则使用系统的Pool
	spawnWorker := func() {
		w = p.workerCache.Get().(*worker)
		w.run()
	}
	// 加锁获取worker
	p.lock.Lock()
	// 弹出一个可用worker
	w = p.workers.detach()
	if w != nil {
		p.lock.Unlock()
	} else if capacity := p.Cap(); capacity == -1 {
		p.lock.Unlock()
		spawnWorker()
	} else if p.Running() < capacity {
		p.lock.Unlock()
		spawnWorker()
	} else {
		if !p.isBlock {
			p.lock.Unlock()
			return
		}
	Reentry:
		p.blockingNum++
		// 等待通知 告诉池中有可用worker
		p.cond.Wait()
		p.blockingNum--
		if p.Running() == 0 {
			p.lock.Unlock()
			spawnWorker()
			return
		}

		w = p.workers.detach()
		if w == nil {
			goto Reentry
		}

		p.lock.Unlock()
	}
	return
}

// revertWorker puts a worker back into free pool, recycling the goroutines.
func (p *Pool) revertWorker(worker *worker) bool {
	if capacity := p.Cap(); (capacity > 0 && p.Running() > capacity) || atomic.LoadInt32(&p.state) == CLOSED {
		return false
	}
	worker.recycleTime = time.Now()
	p.lock.Lock()

	// To avoid memory leaks, add a double check in the lock scope.
	// Issue: https://github.com/panjf2000/ants/issues/113
	if atomic.LoadInt32(&p.state) == CLOSED {
		p.lock.Unlock()
		return false
	}

	err := p.workers.insert(worker)
	if err != nil {
		p.lock.Unlock()
		return false
	}

	// Notify the invoker stuck in 'retrieveWorker()' of there is an available worker in the worker queue.
	p.cond.Signal()
	p.lock.Unlock()
	return true
}

func (p *Pool) getWorker() {

}

// 利用一个协程定期清除过期队列
func (p *Pool) purgePeriodically() {
	expireTicker := time.NewTicker(p.expiryDuration)
	defer expireTicker.Stop()
	// ticker会定期向C的channel发送一条写入数据
	for range expireTicker.C {
		// 若线程池已经关闭 则不在进行过期清理
		if atomic.LoadInt32(&p.state) == CLOSED {
			break
		}
		// 加锁 防止被同时调用
		p.lock.Lock()
		expireWorkers := p.workers.retrieveExpiry(p.expiryDuration)
		p.lock.Unlock()
		for i := range expireWorkers {
			expireWorkers[i].task <- nil
			expireWorkers[i] = nil
		}
		// 协程池运行完毕
		if p.Running() == 0 {
			p.cond.Broadcast()
		}
	}
}

// 池中正在运行的worker数量
func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

// Free returns the available goroutines to work.
func (p *Pool) Free() int {
	return p.Cap() - p.Running()
}

// Cap returns the capacity of this pool.
func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

// Release Closes this pool.
func (p *Pool) Release() {
	atomic.StoreInt32(&p.state, CLOSED)
	p.lock.Lock()
	p.workers.reset()
	p.lock.Unlock()
}

// Reboot reboots a released pool.
func (p *Pool) Reboot() {
	if atomic.CompareAndSwapInt32(&p.state, CLOSED, OPENED) {
		go p.purgePeriodically()
	}
}

// 增加运行数量
func (p *Pool) incRunning() {
	atomic.AddInt32(&p.running, 1)
}

// 减少运行数量
func (p *Pool) decRunning() {
	atomic.AddInt32(&p.running, -1)
}
