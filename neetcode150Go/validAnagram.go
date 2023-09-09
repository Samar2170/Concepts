package main

func validAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smap := make(map[rune]int)
	for _, v := range s {
		smap[v] += 1
	}
	for _, v := range t {
		smap[v] -= 1
	}
	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true
}

// Run time -> 9ms vs 147ms for Python 9/147 = 16.33x faster
