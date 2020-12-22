package pool

import (
	"errors"
	"time"
)

var (
	// errQueueIsFull will be returned when the worker queue is full.
	errQueueIsFull = errors.New("the queue is full")

	// errQueueIsReleased will be returned when trying to insert item to a released worker queue.
	errQueueIsReleased = errors.New("the queue length is zero")
)

type workerArray interface {
	len() int
	isEmpty() bool
	insert(worker *worker) error
	detach() *worker
	retrieveExpiry(duration time.Duration) []*worker
	reset()
}

type arrayType int

func newWorkerArray(size int) workerArray {
	return newWorkerQueue(size)
}
