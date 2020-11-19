package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(i int) (int, iAdder) {
		return i + base, adder2(base + i)
	}
}
func main() {
	f := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, f = f(i)
		fmt.Println(i, s)
	}
}
