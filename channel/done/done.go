package main

import (
	"fmt"
	"sync"
)

// 主题思路是在work接受到任务信息处理完成后向该协程发送处理完成信息，借助通道发送必须有人接收的机制来保证所有的任务均已完成
func doWork(id int, w Worker) {
	go func() {
		for n := range w.data {
			fmt.Printf("Worker %d receive %d\n", id, n)
			w.wg.Done()
		}
	}()
}

// 声明一个结构体
type Worker struct {
	data chan int
	wg   *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		data: make(chan int),
		wg:   wg,
	}
	go doWork(id, w)
	return w
}

// 利用channel 等待任务结束
//func channelDemo() {
//	var workers [10]Worker
//	for id, _ := range workers {
//		workers[id] = createWorker(id)
//	}
//	for id, worker := range workers {
//		worker.data <- id
//	}
//	// 在我们利用通道发送二十次任务后，需要保证20次任务完成后才能结束主线程 所以可以利用另一个通道将完成的信息通知回来
//	// 在将两个通道封装成一个对象，当这个对象的done通道接受到信息后表示该任务完成，但由于通道发送信息必须要有对应的接受操作
//	// 所以再任务发送时由于done操作没有人接受，所以协程会阻塞堵死在循环中，导致第二次任务无法正常进行
//	for id, worker := range workers {
//		worker.data <- id + 10
//	}
//	for _, worker := range workers {
//		<-worker.done
//		<-worker.done
//	}
//}

// 利用waitGroup 等待任务结束
func channelDemo2() {
	var workers [10]Worker
	var wg = sync.WaitGroup{}
	for id, _ := range workers {
		workers[id] = createWorker(id, &wg)
	}
	wg.Add(20)
	for id, worker := range workers {
		worker.data <- id
	}
	for id, worker := range workers {
		worker.data <- id + 10
	}
	wg.Wait()
}
func main() {
	channelDemo2()
}
