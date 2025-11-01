package main

import (
	"fmt"
)

func main() {
	x := 5
	addTen(&x)
	fmt.Println(x) // 输出: 15

	slice := []int{1, 2, 3, 4, 5}
	sliceItemX2(&slice)
	fmt.Printf("切片元素X2: %v\n", slice) // 输出: [2 4 6 8 10]

	RunOddEvenTask()

	RunTaskTimingTask()

	RunShapeTask()
	RunEmployeeTask()

	RunChannelTask()
	RunBufferedChannelTask()

	RunCounterTask()
	RunAtomicCounterTask()
}
