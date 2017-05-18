package main

// 201. Bitwise AND of Numbers Range
// Difficulty: Medium Tags: Bit Manipulation

// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.
// For example, given the range [5, 7], you should return 4.

// this problem is equals to find the left most common bits between m and n
import "fmt"

// iterative
func rangeBitwiseAnd(m, n int) int {
	var factor uint8
	for m < n {
		m >>= 1
		n >>= 1
		factor++
	}
	return m << factor
}

// recursive
func rangeBitwiseAnd1(m, n int) int {
	if n > m {
		return rangeBitwiseAnd1(m>>1, n>>1) << 1
	} else {
		return m
	}
}

// use another trick
func rangeBitwiseAnd2(m, n int) int {
	for n > m {
		n = n & (n - 1)
	}
	return n

}

func main() {
	fmt.Print(rangeBitwiseAnd(5, 7))
	fmt.Print(rangeBitwiseAnd1(5, 7))
	fmt.Print(rangeBitwiseAnd2(5, 7))
}
