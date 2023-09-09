package main

import "fmt"

var Ints = []int{7, 8, 9, 10, 15, 18, 1, 2, 3, 4, 5, 6, 98}

func partition(nums []int, low, high int) int {
	pivot := nums[high]

	for i := low; i < high; i++ {
		if nums[i] > pivot {
			nums[i], nums[high] = nums[high], nums[i]
		}

	}
	fmt.Println(nums)
	return high
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pivot := partition(nums, low, high)
		quickSort(nums, low, pivot-1)
		quickSort(nums, pivot+1, high)
	}
}
