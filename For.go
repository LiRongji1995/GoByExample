package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	/*
	  在go中，range好像是不能用来迭代整数的，这个教程不知道为什么这里用range来迭代整数，这似乎是个错误。
	*/

	//for i := range 3 {
	//	fmt.Println("range", i)
	//}
	for {
		fmt.Println("loop")
		break
	}
	//for n := range 6 {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}
}
