package main

import (
	"fmt"
	"sync"
	"time"
)

func worker2(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}
func main() {
	var wg = sync.WaitGroup{}

	//循环创建 5 个 goroutine。
	for i := 1; i <= 5; i++ {
		/**
		闭包问题修复： 在 for 循环中，为每个 goroutine 创建了
		一个新的变量 i，这样每个 goroutine 都拥有自己的 i 副本，
		避免了多个 goroutine 共享同一个变量的问题。
		在原来的代码中，多个 goroutine 共享同一个变量 i。当
		goroutine 启动时，它们会捕获到 i 的最新值，也就是循环的
		最后一次迭代的值。因此，所有的 goroutine 都认为自己的 id
		是 5，导致输出重复。

		通过为每个 goroutine 创建一个新的变量副本，确保了每个
		goroutine 操作的是不同的变量，从而解决了这个问题。
		*/
		i := i
		wg.Add(1)
		go func() {
			//在函数结束时，调用 wg.Done() 减少 wg 的计数器。
			defer wg.Done()
			worker2(i)
		}()
	}
	//等待 wg 的计数器变为 0，表示所有的 goroutine 都已经完成。
	wg.Wait()
}
