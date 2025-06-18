package main

import "fmt"

// Node represents a node in a circular linked list
type Node struct {
	data int
	next *Node
}

// InsertAtBeginning inserts a node at the beginning of a circular linked list
func InsertAtBeginning(last *Node, data int) *Node {
	// Create a new node
	newNode := &Node{data: data}

	if last == nil {
		newNode.next = newNode
		return newNode
	}

	// Point newNode's next to head (last.next)
	newNode.next = last.next

	// Update last's next to point to new head
	last.next = newNode

	return last
}

// PrintList prints elements of circular linked list
func PrintList(last *Node) {
	if last == nil {
		fmt.Println("List is empty.")
		return
	}

	curr := last.next
	for {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
		if curr == last.next {
			break
		}
	}
	fmt.Println()
}

func main() {
	var last *Node = nil

	// Inserting first node
	last = InsertAtBeginning(last, 30)
	PrintList(last) // Output: 30

	// Inserting more at beginning
	last = InsertAtBeginning(last, 20)
	last = InsertAtBeginning(last, 10)

	PrintList(last) // Output: 10 20 30
}

