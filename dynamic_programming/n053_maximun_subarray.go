package main

// 53. Maximum Subarray
// Difficulty: Medium Tags: Array, Dynamic Programming, Divide and Conquer

// Find the contiguous subarray within an array (containing at least one number) which has the largest sum.

// For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
// the contiguous subarray [4,-1,2,1] has the largest sum = 6.

import "fmt"

// Kadane’s Algorithm:
// Initialize:
// max_so_far = 0
// max_ending_here = 0

// Loop for each element of the array
// (a) max_ending_here = max_ending_here + a[i]
// (b) if(max_ending_here < 0)
// max_ending_here = 0
// (c) if(max_so_far < max_ending_here)
// max_so_far = max_ending_here
// return max_so_far
// note that it doesn't work when all nums are negative
func maxSubArray(nums []int) int {
	max_so_far := 0
	max_ending_here := 0
	// start and end is the index of maxSubArray
	var start, end int
	for i := 0; i < len(nums); i++ {
		max_ending_here = max_ending_here + nums[i]
		if max_ending_here < 0 {
			max_ending_here = 0
			start = i + 1
		}
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
			end = i
		}
	}
	fmt.Println(nums[start : end+1])
	return max_so_far
}

// DP solution: find the sub problem of this
// maxSubArray(A, i) = maxSubArray(A, i - 1) > 0 ? maxSubArray(A, i - 1) : 0 + A[i];
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = 0 + nums[i]
		}
		res = max(dp[i], res)
	}
	return res
}

func main() {
	fmt.Print(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Print(maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
