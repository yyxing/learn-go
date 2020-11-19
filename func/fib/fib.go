package fib

// 测试斐波那契doc
func Fibonacci() func() int {
	// 声明初始变量
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
