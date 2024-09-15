package main

import "fmt"

func main() {
	//创建一个容量为 2 的字符串通道 message。
	//这个通道可以同时存储 2 个字符串。标注了容量就是有缓冲。
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
