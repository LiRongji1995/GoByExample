package main

import "fmt"

/*
*
这段代码定义了一个名为 intSeq 的函数。这个函数本身并不直接返回一个整数，
而是返回另外一个函数。这个被返回的函数的类型是 func() int，也就是说，
它是一个无参数且返回一个整数的函数。
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
}
