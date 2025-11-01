package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type TaskResult struct {
	Index    int
	Duration time.Duration
}

func RunTasksWithTiming(tasks []Task) []TaskResult {
	resultsCh := make(chan TaskResult, len(tasks))
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)
		// 将循环变量作为参数传入，避免闭包捕获问题
		go func(index int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			duration := time.Since(start)
			resultsCh <- TaskResult{Index: index, Duration: duration}
		}(i, task)
	}

	wg.Wait()
	close(resultsCh)

	results := make([]TaskResult, len(tasks))
	for r := range resultsCh {
		results[r.Index] = r
	}
	return results
}

func RunTaskTimingTask() {
	tasks := []Task{
		func() {
			fmt.Printf("任务 0 开始\n")
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("任务 0 结束\n")
		},
		func() {
			fmt.Printf("任务 1 开始\n")
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("任务 1 结束\n")
		},
		func() {
			fmt.Printf("任务 2 开始\n")
			time.Sleep(300 * time.Millisecond)
			fmt.Printf("任务 2 结束\n")
		},
		func() {
			fmt.Printf("任务 3 开始\n")
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("任务 3 结束\n")
		},
		func() {
			fmt.Printf("任务 4 开始\n")
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("任务 4 结束\n")
		},
	}

	results := RunTasksWithTiming(tasks)
	for _, result := range results {
		fmt.Printf("任务 %d 执行时间: %v\n", result.Index, result.Duration)
	}

	fmt.Println("RunTaskTimingTask 完成")

}
