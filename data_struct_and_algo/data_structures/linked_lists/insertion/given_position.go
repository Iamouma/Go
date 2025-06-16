package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Function to insert at a specific position
func insertAtPosition(head *Node, data int, pos int) *Node {
	newNode := &Node{data: data}

	// If inserting at the beginning (position 1)
	if pos == 1 {
		newNode.next = head
		return newNode
	}

	current := head
	count := 1

	// Traverse until the node just before the position
	for current != nil && count < pos-1 {
		current = current.next
		count++
	}

	// If current is nil, position is out of bounds
	if current == nil {
		fmt.Println("Position out of bounds")
		return head
	}

	// Insert the node
	newNode.next = current.next
	current.next = newNode

	return head
}

// Function to print the linked list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	// Create linked list: 3 -> 5 -> 8 -> 10
	node1 := &Node{data: 3}
	node2 := &Node{data: 5}
	node3 := &Node{data: 8}
	node4 := &Node{data: 10}

	node1.next = node2
	node2.next = node3
	node3.next = node4

	// Insert 2 at position 2: 3 -> 2 -> 5 -> 8 -> 10
	head := insertAtPosition(node1, 2, 2)
	printList(head)

	// Insert 11 at position 6: 3 -> 2 -> 5 -> 8 -> 10 -> 11
	head = insertAtPosition(head, 11, 6)
	printList(head)
}

