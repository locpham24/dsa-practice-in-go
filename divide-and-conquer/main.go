package main

import (
	"fmt"
	"math/rand"
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

// Input: 2 4 6 8 10 11 12 13 14, Search: 4
// Output: 1 <=> arr[1] == 4

// Base case: array has 1 item (low == high)
// Recursive case: len(arr) > 0
func BinarySearchRecursion(value int, input []int, low, high int) int {
	if low == high {
		if input[low] == value {
			return low
		}
		return -1
	}

	var median int = (low + high) / 2
	if input[median] == value {
		return median
	} else if input[median] > value {
		return BinarySearchRecursion(value, input, low, median-1)
	} else {
		return BinarySearchRecursion(value, input, median+1, high)
	}
}

func BinarySearch(value int, input []int) int {
	return BinarySearchRecursion(value, input, 0, len(input)-1)
}

// 5, 4, 6, 2, 3, 11, 7, 1, 8
// Input: 5(0), 4(1), 6(2), 2(3), 3(4), 11(5), 7(6), 1(7), 8(8)
// Base case: len(arr) < 2
// Recursive case: len(arr) >= 2 return Recursion(less) + pivot + Recursion(greater)
// Assumsion: pivot = arr[0]
func QuickSortRecursion(input []int) []int {
	if len(input) < 2 {
		return input
	}

	var pivot int = input[0]
	var less, greater []int
	for k, val := range input {
		if val < pivot {
			less = append(less, val)
		} else {
			if k == 0 {
				continue
			}
			greater = append(greater, val)
		}
	}

	result := append(QuickSortRecursion(less), pivot)
	result = append(result, QuickSortRecursion(greater)...)
	return result
}

func QuickSort(input []int) []int {
	return QuickSortRecursion(input)
}

func Swap(input []int, a, b int) []int {
	temp := input[a]
	input[a] = input[b]
	input[b] = temp

	return input
}

// TODO: Should sort in-place, means should not use auxiliary arrays, it will help to reduce space complexity
func RandomizedPartition(arr []int, low, high int) ([]int, int) {
	var pivotIndex int = rand.Intn(high-low) + low
	pivot := arr[pivotIndex]

	var less, greater []int
	for k, val := range arr {
		if val < pivot {
			less = append(less, val)
		} else {
			if k == pivotIndex {
				continue
			}
			greater = append(greater, val)
		}
	}

	result := append(less, pivot)
	result = append(result, greater...)
	return result, pivotIndex
}

func RandomizedQuickSortRecursion(input []int, low, high int) []int {
	if high <= low {
		return input
	}

	input, pivotIndex := RandomizedPartition(input, low, high)

	input = RandomizedQuickSortRecursion(input, low, pivotIndex-1)
	input = RandomizedQuickSortRecursion(input, pivotIndex+1, high)

	return input
}

func RandomizedQuickSort(input []int) []int {
	return RandomizedQuickSortRecursion(input, 0, len(input)-1)
}

// TODO: write test case
func main() {
	fmt.Println("SumOfArrayRecursion: ", SumOfArrayRecursion([]int{1, 2, 3, 4, 5}))
	fmt.Println("BinarySearch: ", BinarySearch(13, []int{2, 4, 6, 8, 10, 11, 12, 13, 14}))
	fmt.Println("QuickSort-FixedPivot-Unsorted-Duplicate: ", QuickSort([]int{5, 4, 6, 2, 3, 2, 11, 7, 1, 8, 2, 2}))
	fmt.Println("QuickSort-FixedPivot-Sorted: ", QuickSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	fmt.Println("QuickSort-PivotRandom-Sorted: ", RandomizedQuickSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	fmt.Println("QuickSort-PivotRandom-Unsorted-Duplicate: ", QuickSort([]int{5, 4, 6, 2, 3, 2, 11, 7, 1, 8, 2, 2}))
}
