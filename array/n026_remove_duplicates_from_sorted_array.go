package main

// 26. Remove Duplicates from Sorted Array
// Difficulty: Easy Tags: Array, Two Pointers
// Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

// Do not allocate extra space for another array, you must do this in place with constant memory.

// For example,
// Given input array nums = [1,1,2],

// Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
// It doesn't matter what you leave beyond the new length.

// removeDuplicates use two-pointers technique two solve this problem.
// Since the array is already sorted, we can keep two pointers ii and jj,
// where ii is the slow-runner while jj is the fast-runner. As long as nums[i] = nums[j]nums[i]=nums[j], we increment jj to skip the duplicate.
// When we encounter nums[j] \neq nums[i]nums[j]≠nums[i], the duplicate run has ended so we must copy its value to nums[i + 1]nums[i+1].
// ii is then incremented and we repeat the same process again until jj reaches the end of array.

import "fmt"

func removeDuplicates(nums []int) int {
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	fmt.Print(removeDuplicates([]int{1, 1, 2, 3, 4}))
}
