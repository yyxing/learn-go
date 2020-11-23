package main

import (
	"fmt"
)

func Worker(id int, c chan int) {
	go func() {
		for n := range c {
			fmt.Printf("Worker %d receive %c\n", id, n)
		}
	}()
}

func CreateWorker(id int) chan<- int {
	c := make(chan int)
	Worker(id, c)
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = CreateWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

}

func channelClose() {
	c := make(chan int)
	go Worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
}
func channelBuffer() {
	var b = make(chan int, 3)
	b <- 3
	b <- 2
	b <- 1
}
func main() {
	//channelDemo()
	//channelBuffer()
	//channelClose()
	var c = make(chan int)
	go func() { c <- 1 }()
	fmt.Println(<-c)
}
