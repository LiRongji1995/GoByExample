package main

/**
这段代码通过互斥锁 (sync.Mutex) 实现线程安全的计数器，
可以用于统计并发场景下的字符串出现次数。
*/
import (
	"fmt"
	"sync"
)

type Container struct {
	//一个 sync.Mutex 类型变量，用于互斥锁，保证并发访问计数器时线程安全
	mu sync.Mutex
	//一个 map[string]int 类型变量，用于存储字符串对应的计数
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}
func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doInCrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	wg.Add(3)
	go doInCrement("a", 10000)
	go doInCrement("a", 10000)
	go doInCrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
