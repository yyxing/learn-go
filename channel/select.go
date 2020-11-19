package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(id int, c chan int) {
	go func() {
		for n := range c {
			time.Sleep(2000 * time.Millisecond)
			fmt.Printf("Worker %d receive %d\n", id, n)
		}
	}()
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	after := time.After(10 * time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-after:
			fmt.Println("程序结束，挤压了", len(values), "数据")
			return
		}
	}

}
