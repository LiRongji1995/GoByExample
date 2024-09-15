package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
*
这段代码演示了如何使用 atomic.Uint64 和 sync.WaitGroup
来实现并发计数，确保在多个 goroutine 并发访问计数器时不会出现
数据竞争，并且可以正确地统计总的计数结果。
*/
func main() {
	var ops atomic.Uint64
	//用于同步多个 goroutine 的执行。
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())
}
