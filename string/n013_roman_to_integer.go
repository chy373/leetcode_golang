package main

// 13. Roman to Integer
// Difficulty: Easy Tags: Math, String

// Given a roman numeral, convert it to an integer.
// Input is guaranteed to be within the range from 1 to 3999.

import "fmt"

//roman -- int I(1)，V(5)，X(10)，L(50)，C(100)，D(500)，M(1000)
// V L D can't be substracted
func romanToInt(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		switch string(s[i]) {
		case "I":
			if res < 5 {
				res += 1
			} else {
				res += -1
			}
			break
		case "V":
			res += 5
			break
		case "X":
			if res < 50 {
				res += 10
			} else {
				res += -10
			}
			break
		case "L":
			res += 50
			break
		case "C":
			if res < 500 {
				res += 100
			} else {
				res += -100
			}
			break
		case "D":
			res += 500
			break
		case "M":
			res += 1000
			break
		}
	}
	return res
}

func main() {
	fmt.Print(romanToInt("XII"))
}
