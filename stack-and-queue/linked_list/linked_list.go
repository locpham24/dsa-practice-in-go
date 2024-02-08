package main

import (
	"fmt"
)

type LinkedList struct {
	First *Node
}

type Node struct {
	Value int
	Next *Node
}

func (l *LinkedList) isEmpty() bool {
	return l.First == nil
}

func (l *LinkedList) push(value int) {
	newNode := &Node{
		Value: value,
		Next:  l.First,
	}
	l.First = newNode
}

func (l *LinkedList) pop() {
	fmt.Printf("%d \n", l.First.Value)
	l.First = l.First.Next
}

func main() {
	l := LinkedList{}
	l.push(1)
	l.push(2)
	l.push(5)
	l.pop()
	l.push(3)
	l.push(4)
	l.pop()
	l.pop()
	l.pop()
	l.pop()

	// expect: 5 4 3 2 1
}