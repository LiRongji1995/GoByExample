package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTimer(500 * time.Millisecond)
	//创建一个布尔通道 done，用于通知 goroutine 停止。
	done := make(chan bool)

	go func() {
		for {
			select {
			//如果 done 通道有数据可读，则退出循环。
			case <-done:
				return
			case t := <-ticker.C: //如果 ticker.C 通道有数据可读，则打印时间。
				fmt.Println("Tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	//向 done 通道发送 true，通知 goroutine 停止。
	done <- true
	fmt.Println("Ticker stopped")
}
