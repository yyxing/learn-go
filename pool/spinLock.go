package pool

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock uint32

// 自旋加锁
func (sp *spinLock) Lock() {
	// 利用CAS不断将当前标识的0变成1 若更改成功表示加锁成功 否则让出cpu
	for !atomic.CompareAndSwapUint32((*uint32)(sp), 0, 1) {
		runtime.Gosched()
	}
}

// 自旋解锁 将标识为更改成0
func (sp *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sp), 0)
}
func NewSpinLock() sync.Locker {
	return new(spinLock)
}
