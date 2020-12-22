package pool

import "time"

type queue struct {
	items  []*worker
	expiry []*worker
	head   int
	tail   int
	size   int
	isFull bool
}

func newWorkerQueue(size int) *queue {
	return &queue{
		size:  size,
		items: make([]*worker, size),
	}
}

// 判断队列的长度
func (wq *queue) len() int {
	if wq.size == 0 {
		return 0
	}
	// 若头尾一致则存在满空两种情况
	if wq.head == wq.tail {
		if wq.isFull {
			return wq.size
		}
		return 0
	}
	// 队列未满时
	if wq.tail > wq.head {
		return wq.tail - wq.head
	}
	// 队列已经开始循环了
	return wq.size - wq.head + wq.tail
}

// 队列是否为空
func (wq *queue) isEmpty() bool {
	return wq.head == wq.tail && !wq.isFull
}

// 插入数据到队列中
func (wq *queue) insert(worker *worker) error {
	if wq.size == 0 {
		return errQueueIsReleased
	}

	if wq.isFull {
		return errQueueIsFull
	}
	wq.items[wq.tail] = worker
	wq.tail++
	if wq.tail == wq.size {
		wq.tail = 0
	}
	if wq.head == wq.tail {
		wq.isFull = true
	}
	return nil
}

// 弹出第一个元素
func (wq *queue) detach() *worker {
	if wq.isEmpty() {
		return nil
	}
	w := wq.items[wq.head]
	wq.items[wq.head] = nil
	wq.head++
	if wq.head == wq.size {
		wq.head = 0
	}
	return w
}

// 清除过期worker
func (wq *queue) retrieveExpiry(duration time.Duration) []*worker {
	if wq.isEmpty() {
		return nil
	}
	// 清空过期队列
	wq.expiry = wq.expiry[:0]
	// 过期时间
	expiryTime := time.Now().Add(-duration)

	for !wq.isEmpty() {
		// 表示没过期 由于队列按照插入时间排列 所以直接break
		if expiryTime.Before(wq.items[wq.head].recycleTime) {
			break
		}
		// 超时 加入超时队列 将可用队列中的超时worker清除
		wq.expiry = append(wq.expiry, wq.items[wq.head])
		wq.items[wq.head] = nil
		wq.head++
		if wq.head == wq.size {
			wq.head = 0
		}
		wq.isFull = false
	}

	return wq.expiry
}

// 重置等待队列
func (wq *queue) reset() {
	if wq.isEmpty() {
		return
	}
Releasing:
	if w := wq.detach(); w != nil {
		w.task <- nil
		goto Releasing
	}
	wq.items = wq.items[:0]
	wq.head = 0
	wq.size = 0
	wq.size = 0
}
