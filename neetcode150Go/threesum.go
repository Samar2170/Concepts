package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var i int
	var result [][]int
	for i <= len(nums) {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
		i++
	}
	return result
}
