package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	testcases := []struct {
		arr []int
	}{
		{[]int{1, 2, 3, 1}},
		{[]int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
	}

	for _, testcase := range testcases {
		fmt.Printf("containsDuplicate(%v) = %v\n", testcase.arr, containsDuplicate(testcase.arr))
	}
}
