package main

import (
	"fmt"
)

// Input: 2 4 6
// Output: 12 = 2 + 4 + 6

// Base case: len(arr) == 2
// Recursive case: len(arr) > 2
func SumOfArrayRecursion(input []int) int {
	length := len(input)
	if length == 2 {
		return input[0] + input[1]
	}
	return SumOfArrayRecursion(input[:length-1]) + input[length-1]
}

func main() {
	sum := SumOfArrayRecursion([]int{1, 2, 3, 4, 5})
	fmt.Println("sum: ", sum)
}
