package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func lengthOfLongestSubstring(s string) int {
	i, j := 0, 1
	maxlen := 0
	chars := make(map[string]int)
	for j < len(s) {
		fmt.Println(chars)
		if _, ok := chars[string(s[j])]; ok {
			maxlen = max(j-1-i, maxlen)
			i += 1
			j = i + 1
		} else {
			chars[string(s[j])] = j
			chars[string(s[i])] = i
			j += 1
		}
	}
	return maxlen
}

func main() {
	fmt.Print(lengthOfLongestSubstring("pwwkew"))
}
