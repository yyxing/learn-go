package pool

import (
	"github.com/sirupsen/logrus"
	"time"
)

type worker struct {
	pool *Pool
	// 任务通道
	task chan func()
	// 上次运行时间 用于排序
	recycleTime time.Time
}

func (w *worker) run() {
	w.pool.incRunning()
	go func() {
		defer func() {
			w.pool.decRunning()
			w.pool.workerCache.Put(w)
			if p := recover(); p != nil {
				logrus.Info(p)
			}
		}()
		for f := range w.task {
			if f == nil {
				return
			}
			f()
			if ok := w.pool.revertWorker(w); !ok {
				return
			}
		}
	}()
}
