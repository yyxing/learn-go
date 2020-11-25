package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
)

// 函数实现接口
type intGen func() int

func (gen intGen) Read(p []byte) (n int, err error) {
	next := gen()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	//var f intGen = fib.Fibonacci()
	//printFileContents(f)
	var wp = sync.WaitGroup{}
	wp.Add(100)
	sum := 0
	sumTemp := 0
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				sumTemp = sumTemp + 1
				sum = sumTemp
			}
			wp.Done()
		}()
	}
	wp.Wait()
	fmt.Println(sumTemp == 1000)
	fmt.Println(sumTemp)
}
