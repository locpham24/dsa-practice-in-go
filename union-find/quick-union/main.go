package main

import "fmt"

// Weighted Quick Union
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sz := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	union(arr, sz, 4, 3)
	union(arr, sz, 3, 8)
	union(arr, sz, 6, 5)
	union(arr, sz, 9, 4)
	union(arr, sz, 2, 1)
	union(arr, sz, 5, 0)
	union(arr, sz, 7, 2)
	union(arr, sz, 6, 1)
	union(arr, sz, 7, 3)
}

func root(arr []int, a int) int {
	for a != arr[a] {
		a = arr[a]
	}
	return a
}

// Assume array is not empty
func find(arr []int, a, b int) bool {
	return root(arr, a) == root(arr, b)
}

// Assume array is not empty
func union(arr []int, sz []int, a, b int) []int {
	rootA := root(arr, a)
	rootB := root(arr, b)

	if rootA == rootB return
	
	if sz[rootA] < sz[rootB] {
		arr[rootA] = rootB
		sz[rootB] += sz[rootA]
	} else {
		arr[rootB] = rootA
		sz[rootA] += sz[rootB]
	}

	fmt.Printf("union[%d-%d]: %v \n",a, b, arr)
	return arr
}