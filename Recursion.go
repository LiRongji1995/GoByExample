package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println(fact(2))

	var fib func(n int) int //声明了一个名为fib的变量，这个变量可以存储一个接受一个整数作为参数，并且返回一个整数的函数。

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
