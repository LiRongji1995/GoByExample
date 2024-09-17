package main

import "fmt"

func mayPanic() {
	panic("a problem")
}
func main() {
	//在 main 函数中使用 defer 关键字注册了一个匿名函数，这个函数会在 main 函数返回
	//或发生 panic 时执行。
	defer func() {
		//如果 recover 捕获到了 panic，就打印出错误信息，然后程序继续执行。
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic")
}
