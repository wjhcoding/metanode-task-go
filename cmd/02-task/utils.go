package main

import (
	"fmt"
	"sync"
	"time"
)

func addTen(p *int) {
	*p += 10
}

func sliceItemX2(p *[]int) {
	s := *p
	for i := range s {
		s[i] *= 2
	}
}

func RunOddEvenTask() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 打印奇数的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			// 增加日志输出
			fmt.Printf("协程打印奇数: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 打印偶数的协程
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("协程打印偶数: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	wg.Wait()

	fmt.Println("RunOddEvenTask 完成")

}
