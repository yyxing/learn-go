package fib

import "fmt"

func ExampleFibonacci() {
	f := Fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	//Output:
	//1
	//1
	//2
	//3
	//5
}
