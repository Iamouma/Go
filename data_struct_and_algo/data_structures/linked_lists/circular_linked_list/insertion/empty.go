package main

import "fmt"

// Node structure in a circular linked list
type Node struct {
	data int
	next *Node
}

// Inserts a node into an empty circular linked list
func InsertInEmpty(last *Node, data int) *Node {
	if last != nil {
		return last
	}

	// create a new node
	newNode := &Node{data: data}
	newNode.next = newNode // Point to itself

	last = newNode
	return last
}

// prints elements of circular linked list
func PrintList(last *Node) {
	if last == nil {
		fmt.Println("List is empty.")
		return
	}

	first := last.next
	curr := first
	for {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
		if curr == first {
			break
		}
	}
	fmt.Println()
}

func main() {
	var last *Node = nil

	last = InsertInEmpty(last, 10)

	fmt.Print("Circular Linked List: ")
	PrintList(last)
}
