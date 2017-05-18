package main

// 3. Longest Substring Without Repeating Characters
// Difficulty: Medium Tags: Hash Table, Two Pointers, String
// Given a string, find the length of the longest substring without repeating characters.

// Examples:

// Given "abcabcbb", the answer is "abc", which the length is 3.

// Given "bbbbb", the answer is "b", with the length of 1.

// Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	ref := make(map[rune]int)
	var start, longest int = 0, 0
	for i, char := range s {
		if idx, ok := ref[char]; ok {
			start = max(start, idx+1)
		}
		ref[char] = i
		longest = max(longest, i-start+1)
	}
	return longest
}

func main() {
	fmt.Print(lengthOfLongestSubstring("pwwkew"))
}
