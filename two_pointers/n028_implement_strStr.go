package main

// 28. Implement strStr()
// Difficulty: Easy Tags: String, Two Pointers

// Implement strStr().

// Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
import "fmt"

// brute force
func strStr(haystack string, needle string) int {
	m, n := len(needle), len(haystack)
	if m == 0 {
		return 0
	}
	for i := 0; i < n-m+1; i++ {
		for j := 0; j < m; j++ {
			if needle[j] != haystack[i+j] {
				break
			}
			if j+1 == m {
				return i
			}
		}
	}
	return -1
}

// elegent brute force
func strStr1(s string, t string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(t) {
				return i
			}
			if i+j == len(s) {
				return -1
			}
			if t[j] != s[i+j] {
				break
			}
		}
	}
}

// todo: KMP

func main() {
	fmt.Print(strStr("mississippi", "issip"))
	fmt.Print(strStr1("mississippi", "issip"))
}
