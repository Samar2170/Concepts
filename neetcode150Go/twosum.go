package main

func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)
	var result []int
	for i, v := range nums {
		if _, ok := nmap[v]; ok {
			result = append(result, nmap[v])
			result = append(result, i)
			return result
		} else {
			nmap[target-v] = i
		}
	}
	return result
}

// Runtime -> 3ms vs 138ms for Python 138/3 = 46.00x faster
