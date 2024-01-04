package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sz := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	expected := []int{6, 2, 6, 4, 6, 6, 6, 2, 4, 4}
	
	union(arr, sz, 4, 3)
	union(arr, sz, 3, 8)
	union(arr, sz, 6, 5)
	union(arr, sz, 9, 4)
	union(arr, sz, 2, 1)
	union(arr, sz, 5, 0)
	union(arr, sz, 7, 2)
	union(arr, sz, 6, 1)
	union(arr, sz, 7, 3)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Test failed. Expected %v, got %v", expected, arr)
	}
}