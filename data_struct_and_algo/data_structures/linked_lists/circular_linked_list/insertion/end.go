package main

import "fmt"

// Node represents a node in a circular linked list
type Node struct {
	data int
	next *Node
}

// InsertAtEnd inserts a node at the end of a circular linked list
func InsertAtEnd(last *Node, data int) *Node {
	newNode := &Node{data: data}

	if last == nil {
		newNode.next = newNode
		return newNode
	}

	newNode.next = last.next // head
	last.next = newNode
	last = newNode // new last node

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

	// Inserting at end into an empty list
	last = InsertAtEnd(last, 10)
	PrintList(last) // Output: 10

	// Inserting more at end
	last = InsertAtEnd(last, 20)
	last = InsertAtEnd(last, 30)
	last = InsertAtEnd(last, 40)

	PrintList(last) // Output: 10 20 30 40
}

