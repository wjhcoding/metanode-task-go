package main

import (
	"reflect"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		input  []int
		expect int
	}{
		{[]int{2, 3, 2, 4, 4}, 3},
		{[]int{1, 1, 2}, 2},
		{[]int{7}, 7},
	}
	for _, c := range cases {
		if got := singleNumber(c.input); got != c.expect {
			t.Errorf("singleNumber(%v) = %d, want %d", c.input, got, c.expect)
		}
	}
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"()[]{}", true},
		{"([)]", false},
		{"([])", true},
		{"((({[]})))", true},
		{"((({[}))", false},
	}
	for _, c := range cases {
		if got := isValid(c.input); got != c.expect {
			t.Errorf("isValid(%q) = %v, want %v", c.input, got, c.expect)
		}
	}
}

func TestPlusOne(t *testing.T) {
	cases := []struct {
		input  []int
		expect []int
	}{
		{[]int{9}, []int{1, 0}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{0}, []int{1}},
	}
	for _, c := range cases {
		got := plusOne(append([]int{}, c.input...))
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("plusOne(%v) = %v, want %v", c.input, got, c.expect)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		input   []int
		expect  []int
		expectK int
	}{
		{[]int{1, 1, 2, 2, 3, 3, 3, 4}, []int{1, 2, 3, 4}, 4},
		{[]int{1, 1, 1, 1}, []int{1}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 4},
		{[]int{}, []int{}, 0},
	}
	for _, c := range cases {
		nums := append([]int{}, c.input...)
		k := removeDuplicates(nums)
		if k != c.expectK {
			t.Errorf("removeDuplicates(%v) k = %d, want %d", c.input, k, c.expectK)
		}
		if !reflect.DeepEqual(nums[:k], c.expect) {
			t.Errorf("removeDuplicates(%v) result = %v, want %v", c.input, nums[:k], c.expect)
		}
	}
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		expect []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for _, c := range cases {
		got := twoSum(c.nums, c.target)
		if !reflect.DeepEqual(got, c.expect) && !reflect.DeepEqual(got, []int{c.expect[1], c.expect[0]}) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", c.nums, c.target, got, c.expect)
		}
	}
}
