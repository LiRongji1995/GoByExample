package main

import "fmt"

func main() {
	// 创建一个无缓冲的字符串通道
	message := make(chan string)

	go func() { message <- "ping" }()

	//由于 message 是无缓冲的，这个接收操作也会阻塞，
	//直到有 goroutine 向这个通道发送数据。
	msg := <-message
	fmt.Println(msg)
}
