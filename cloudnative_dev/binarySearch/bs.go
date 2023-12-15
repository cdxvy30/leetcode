package main

import "fmt"

func BSNonDupNum(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Check if the mid index is even or odd.
		if mid%2 == 0 {
			// If the mid index is even, its duplicate should be on its right (if it has a duplicate).
			if nums[mid] == nums[mid+1] {
				left = mid + 2
			} else {
				right = mid
			}
		} else {
			// If the mid index is odd, its duplicate should be on its left (if it has a duplicate).
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return nums[left]
}

func main() {
	nums := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 8, 8, 10}
	result := BSNonDupNum(nums)
	fmt.Println("The single element is:", result)
}
