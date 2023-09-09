package main

// contains

func containsDuplicates(nums []int) bool {
	nmap := make(map[int]int)
	for _, num := range nums {
		if _, ok := nmap[num]; ok {
			nmap[num] += 1
		} else {
			nmap[num] = 1
		}
	}

	for _, v := range nmap {
		if v > 1 {
			return true
		}
	}
	return false
}

// Run time -> 83ms vs 486ms for Python 486/83 =  5.85x faster
