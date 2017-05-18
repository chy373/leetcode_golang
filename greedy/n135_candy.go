package main

// 135. Candy
// Difficulty: Hard Tags: Greedy

// There are N children standing in a line. Each child is assigned a rating value.
// You are giving candies to these children subjected to the following requirements:
// 1. Each child must have at least one candy.
// 2 . Children with a higher rating get more candies than their neighbors.
// What is the minimum candies you must give?

// there is another solution with O(1) space O(n) time
// ref: http://www.allenlipeng47.com/blog/index.php/2016/07/21/candy/

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func candy(rating []int) int {
	size := len(rating)
	var candies = make([]int, size)
	// init a slice with 1
	for i := 0; i < size; i++ {
		candies[i] = 1
	}
	// make sure every higher rate child have 1 more candy than his left neighbor
	for i := 1; i < size; i++ {
		if rating[i] > rating[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	// make sure every higher rate child have the max of (candies[i-1]+1, candies[i])
	for i := size - 2; i >= 0; i-- {
		if rating[i] > rating[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}
	// get the total candies
	total := 0
	for _, num := range candies {
		total += num
	}
	return total
}

func main() {
	fmt.Print(candy([]int{23, 1}))
}
