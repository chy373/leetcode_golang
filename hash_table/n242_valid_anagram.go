package main

// 242. Valid Anagram
// difficulty: Easy tags: Hash Table, Sort

// Given two strings s and t, write a function to determine if t is an anagram of s.

// For example,
// s = "anagram", t = "nagaram", return true.
// s = "rat", t = "car", return false.
import "fmt"
import "reflect"

func isAnagram(s string, t string) bool {
	a, b := make(map[rune]int), make(map[rune]int)
	for _, char := range s {
		a[char] += 1
	}
	for _, char := range t {
		b[char] += 1
	}
	return reflect.DeepEqual(a, b)
}

func isAnagram1(s string, t string) bool {
	ref := [26]int{}
	for _, char := range s {
		ref[char-'a']++
	}
	for _, char := range t {
		ref[char-'a']--
	}
	for _, n := range ref {
		if n != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(isAnagram("abc", "cba"))
	fmt.Print(isAnagram1("a", "b"))
}
