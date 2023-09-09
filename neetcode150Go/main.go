package main

func main() {
	// transact()
	// testTKV()
	testIt()
}

func testQuicksort() {
	quickSort(Ints, 0, len(Ints)-1)
}

func testThreeSum() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	println(threeSum(nums))
}

func testContainsDups() {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	println(containsDuplicates(nums1))
	println(containsDuplicates(nums2))
	println(containsDuplicates(nums3))

}

func testValidAnagram() {
	println(validAnagram("anagram", "nagaram"))
	println(validAnagram("rat", "car"))
}

func testTwosum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	println(twoSum(nums, target))
	nums2 := []int{3, 2, 4}
	target2 := 6
	println(twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	println(twoSum(nums3, target3))
}
