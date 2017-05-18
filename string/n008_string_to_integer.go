package main

// 8. String to Integer (atoi)
// Difficulty: Easy Tags: Math, String
// Implement atoi to convert a string to an integer.

// Hint: Carefully consider all possible input cases. If you want a challenge, please do not see below and ask yourself what are the possible input cases.

// Notes: It is intended for this problem to be specified vaguely (ie, no given input specs). You are responsible to gather all the input requirements up front.

// Update (2015-02-10):
// The signature of the C++ function had been updated. If you still see your function signature accepts a const char * argument, please click the reload button  to reset your code definition.

// Requirements for atoi:
// The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

// The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

// If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

// If no valid conversion could be performed, a zero value is returned. If the correct value is out of the range of representable values, INT_MAX (2147483647) or INT_MIN (-2147483648) is returned.

// handle four cases
// 1. invalid input(empty string)
// 2. white leading space
// 3. sign
// 4. overflow
// 5. break when occurs non-digit
import "fmt"

const (
	MaxUint = ^uint32(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func myAtoi(str string) int {

	idx, sign, total := 0, 1, 0
	// 1. empty string
	if len(str) == 0 {
		return 0
	}
	// 2. remove leading space
	for str[idx] == ' ' {
		idx++
	}
	// 3. sign
	if str[idx] == '+' {
		sign = 1
		idx++
	} else if str[idx] == '-' {
		sign = -1
		idx++
	}
	for idx < len(str) {
		digit := int(str[idx] - '0')
		// 4. invalid digit
		if digit < 0 || digit > 9 {
			break
		}
		// 5. take care of overflow
		if MaxInt/10 < total || MaxInt/10 == total && MaxInt%10 < digit {
			if sign == -1 {
				return MinInt
			} else {
				return MaxInt
			}
		}
		total = total*10 + digit
		idx++
	}
	return sign * total
}

func main() {
	fmt.Println(myAtoi("1"))
}
