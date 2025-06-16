package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// function to delete the last node
func deleteAtEnd(head *Node) *Node {
	// check if list is empty
	if head == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// check if only one node exists
	if head.next == nil {
		return nil
	}

	// traverse to find the second last one
	current := head
	for current.next.next != nil {
		current = current.next
	}

	// delete the last node
	current.next = nil

	// return the updated head
	return head
}

// Function to print the list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

func main() {
	// Create linked list: 2 -> 4 -> 6 -> 8
	node1 := &Node{data: 2}
	node2 := &Node{data: 4}
	node3 := &Node{data: 6}
	node4 := &Node{data: 8}

	node1.next = node2
	node2.next = node3
	node3.next = node4

	head := node1

	fmt.Print("Original List: ")
	printList(head)

	// Delete last node
	head = deleteAtEnd(head)

	fmt.Print("After Deletion: ")
	printList(head)
}
