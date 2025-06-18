package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// DeleteFirstNode removes the first node from circular linked list
func DeleteFirstNode(last *Node) *Node {
	if last == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// Only one node
	if last.next == last {
		return nil
	}

	// More than one node
	head := last.next
	last.next = head.next
	head = nil // optional in Go (GC will reclaim it)

	return last
}

// Helper function to print circular linked list
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

// InsertAtEnd inserts a new node at the end of a circular linked list
func InsertAtEnd(last *Node, data int) *Node {
	newNode := &Node{data: data}
	if last == nil {
		newNode.next = newNode
		return newNode
	}
	newNode.next = last.next
	last.next = newNode
	return newNode
}

func main() {
	var last *Node

	last = InsertAtEnd(last, 10)
	last = InsertAtEnd(last, 20)
	last = InsertAtEnd(last, 30)

	fmt.Print("Original List: ")
	PrintList(last)

	last = DeleteFirstNode(last)

	fmt.Print("After Deleting First Node: ")
	PrintList(last)
}

