package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用原子操作实现的无锁计数器
type AtomicCounter struct {
	value int64
}

// 原子递增
func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// 原子读取
func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

// 使用 sync/atomic 的并发安全计数器
func RunAtomicCounterTask() {
	var wg sync.WaitGroup
	counter := &AtomicCounter{}

	numGoroutines := 10
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数值: %d（预期 %d）\n", counter.Value(), numGoroutines*incrementsPerGoroutine)
}
