package main

// 217. Contains Duplicate
// Difficulty: Easy Tags: Array HashTable

// Given an array of integers, find if the array contains any duplicates.
// Your function should return true if any value appears at least twice in the array,
// and it should return false if every element is distinct.
import "fmt"

func containsDuplicate(nums []int) bool {
	var ref = map[int]bool{}
	for _, num := range nums {
		if ref[num] {
			return true
		} else {
			ref[num] = true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 5, 4}
	fmt.Print(containsDuplicate(nums))
}
