package main

import "fmt"
// 0 1 2 3 4 5 6 7
// 1 1 3 3 1 3 1 2

func main() {
	// fmt.Println("test")
	arr := []int{1, 1, 3, 3, 1, 3, 1, 2}
	fmt.Println("isconnect", IsConnected(arr, 2, 6))

	fmt.Println("connect", Connect(arr, 2, 6))

	fmt.Println("isconnect", IsConnected(arr, 2, 6))
}

// Assume array is not empty
func IsConnected(arr []int, a, b int) bool {
	return arr[a] == arr[b]
}

// Assume array is not empty
func Connect(arr []int, a, b int) []int {
	i := 0
	value := arr[a]
	for i < len(arr) - 1 {	
		if arr[i] == value {
			arr[i] = arr[b]
		}
		i++
	}
	return arr
}