package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 2, 4, 4}
	fmt.Println("只出现一次的元素:", singleNumber(nums))
	fmt.Println("括号有效性:", isValid("()[]{}"))       // true
	fmt.Println("括号有效性:", isValid("([)]"))         // false
	fmt.Println("括号有效性:", isValid("([])"))         // false
	fmt.Println("大整数加一:", plusOne([]int{9}))       // [1 0]
	fmt.Println("大整数加一:", plusOne([]int{9, 9, 9})) // [1 0 0 0]

	// removeDuplicates 示例
	nums2 := []int{1, 1, 2, 2, 3, 3, 3, 4}
	k := removeDuplicates(nums2)
	fmt.Printf("去重后唯一元素个数: %d, 数组: %v\n", k, nums2[:k])

	fmt.Println("两数之和下标:", twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	counts := make(map[int]int)
	for _, x := range nums {
		counts[x]++
	}
	for k, v := range counts {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 判断括号字符串是否有效
func isValid(s string) bool {
	stack := []rune{}
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 大整数加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}
	// 最高位有进位
	return append([]int{1}, digits...)
}

// 原地去重，返回唯一元素个数
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

// 两数之和，返回下标
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
