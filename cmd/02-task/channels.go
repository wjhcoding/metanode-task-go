package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者：向通道中发送 1~10 的整数
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch) // 所有数据发送完毕后关闭通道
}

// 生产者：发送 100 个整数
func producerHundred(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("发送: %d（通道中剩余 %d 个）\n", i, len(ch))
		time.Sleep(10 * time.Millisecond)
	}
	close(ch)
}

// 消费者：从通道中接收数据并打印
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("接收: %d\n", num)
	}
}

// 消费者：接收并打印整数
func consumerHundred(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("接收: %d（通道中剩余 %d 个）\n", num, len(ch))
		time.Sleep(20 * time.Millisecond)
	}
}

// 运行无缓冲通道任务
func RunChannelTask() {
	var wg sync.WaitGroup

	messageCh := make(chan int)

	wg.Add(2)
	go producer(messageCh, &wg) // 启动生产者和消费者协程
	go consumer(messageCh, &wg)

	wg.Wait()
	fmt.Println("RunChannelTask 完成")
}

// 运行带缓冲通道
func RunBufferedChannelTask() {
	var wg sync.WaitGroup

	bufferedCh := make(chan int, 10) // 创建一个容量为10的缓冲通道

	wg.Add(2)
	go producerHundred(bufferedCh, &wg)
	go consumerHundred(bufferedCh, &wg)

	wg.Wait()
	fmt.Println("RunBufferedChannelTask 完成")
}
