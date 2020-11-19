package main

import (
	"awesomeProject/func/fib"
	"bufio"
	"fmt"
	"io"
	"strings"
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
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
