package main

import (
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	// Test case 1: Basic test with valid inputs
	arr1 := []int{1, 2, 3, 4, 5}
	result1 := Connect(arr1, 2, 3)
	expected1 := []int{1, 2, 4, 4, 5}
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected1, result1)
	}

	// Test case 2: Array with all elements equal
	arr2 := []int{3, 3, 3, 3, 3}
	result2 := Connect(arr2, 0, 4)
	expected2 := []int{3, 3, 3, 3, 3}
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected2, result2)
	}
}

func TestConnectAndIsConnected(t *testing.T) {
	arr := []int{1, 1, 3, 3, 1, 3, 1, 2}

	if IsConnected(arr, 1, 6) != true {
		t.Errorf("Test case 1 failed. Expected %v", true)
	}

	if IsConnected(arr, 1, 5) != false {
		t.Errorf("Test case 2 failed. Expected %v", false)
	}

	_ = Connect(arr, 1, 2)
	expected := []int{3, 3, 3, 3, 3, 3, 3, 2}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected, arr)
	}

	if IsConnected(arr, 1, 2) != true {
		t.Errorf("Test case 4 failed. Expected %v", true)
	}

	if IsConnected(arr, 1, 5) != true {
		t.Errorf("Test case 5 failed. Expected %v", true)
	}
}