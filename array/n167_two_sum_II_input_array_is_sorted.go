package main

// 167. Two Sum II - Input array is sorted
// difficulty: Medium tags: Array, Two Pointers, Binary Search
// Given an array of integers that is already sorted in ascending order,
// find two numbers such that they add up to a specific target number.

// The function twoSum should return indices of the two numbers such that they add up to the target,
// where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

// You may assume that each input would have exactly one solution.

// Input: numbers={2, 7, 11, 15}, target=9
// Output: index1=1, index2=2

import "fmt"

func twoSum(numbers []int, target int) []int {
	var i, j int = 0, len(numbers) - 1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}

func main() {
	fmt.Print(twoSum([]int{1, 2, 3, 5}, 4))
}
