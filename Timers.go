package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置超时时间为 2 秒。
	timer1 := time.NewTimer(2 * time.Second)
	// 等待定时器 timer1 超时
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	//goroutine 阻塞在 <-timer2.C:
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	//所以先尝试停止定时器 timer2
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
