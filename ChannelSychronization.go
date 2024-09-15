package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
func main() {
	done := make(chan bool, 1)
	go worker(done)

	/**
	作用一，从 done 通道接收一个值：这表示主 goroutine 正在等待 worker goroutine
	完成任务并发送信号到 done 通道。作用二，阻塞主 goroutine：
	直到 worker goroutine 发送信号，主 goroutine 才会继续执行。
	如果没有 <-done，主 goroutine 不会再等待 worker goroutine 完成。
	*/
	<-done
}
