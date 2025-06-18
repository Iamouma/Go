package main

import (
	"fmt"
)

// Node defines a node in a doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// InsertAtPosition inserts a new node with the given data at the given position
func InsertAtPosition(head *Node, data int, position int) *Node {
	newNode := &Node{data: data}

	// Insertion at position 1
	if position == 1 {
		newNode.next = head
		if head != nil {
			head.prev = newNode
		}
		return newNode
	}

	curr := head
	count := 1

	// Traverse to position - 1
	for curr != nil && count < position-1 {
		curr = curr.next
		count++
	}

	// If position is invalid (out of range)
	if curr == nil {
		fmt.Println("Position out of bounds.")
		return head
	}

	// Insert newNode after curr
	newNode.next = curr.next
	newNode.prev = curr
	curr.next = newNode

	if newNode.next != nil {
		newNode.next.prev = newNode
	}

	return head
}

// PrintList prints the doubly linked list
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.data)
		head = head.next
	}
	fmt.Println()
}

func main() {
	var head *Node

	head = InsertAtPosition(head, 10, 1) // 10
	head = InsertAtPosition(head, 20, 2) // 10 -> 20
	head = InsertAtPosition(head, 30, 3) // 10 -> 20 -> 30
	head = InsertAtPosition(head, 25, 3) // 10 -> 20 -> 25 -> 30
	head = InsertAtPosition(head, 5, 1)  // 5 -> 10 -> 20 -> 25 -> 30

	fmt.Print("Doubly Linked List: ")
	PrintList(head)
}

