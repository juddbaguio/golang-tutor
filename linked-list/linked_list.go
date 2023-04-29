package main

import "fmt"

type Node struct {
	Value int
	Next  *Node // memory address of next Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *LinkedList) Insert(data int) {
	// initial case
	if ll.Head == nil && ll.Tail == nil {
		ll.Head = &Node{
			Value: data,
		}

		ll.Tail = ll.Head
		// at first initialization, Head and Tail have the same pointer
		return
	}

	ll.Tail.Next = &Node{
		Value: data,
	}

	ll.Tail = ll.Tail.Next
}

func (ll *LinkedList) Traverse() {
	currentNode := ll.Head
	for currentNode != nil {
		fmt.Println(currentNode.Value)
		currentNode = currentNode.Next
	}
}

func (ll *LinkedList) DeleteAt(idx int) {
	if ll.Head == nil && ll.Tail == nil {
		panic("index out of bounds")
	}

	currIndex := 0
	currentNode := ll.Head
	var previousNode *Node

	for currentNode != nil {
		if currIndex == idx {
			nextNode := currentNode.Next
			previousNode.Next = nextNode
			return
		}
		currIndex += 1
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	panic("index out of bounds")
}
