package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Function to modify node at a given position (1-based index)
func modifyAtPosition(head *Node, position int, newValue int) *Node {
	// Step 1: Check for empty list
	if head == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// Step 2: Traverse to the required position
	current := head
	for i := 1; i < position && current != nil; i++ {
		current = current.next
	}

	// Step 3: Check for valid position
	if current == nil {
		fmt.Println("Position out of bounds.")
		return head
	}

	// Step 4: Update the node's data
	current.data = newValue
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
	// Create linked list: 8 -> 2 -> 3 -> 1 -> 7
	node1 := &Node{data: 8}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 1}
	node5 := &Node{data: 7}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	head := node1

	fmt.Print("Original List: ")
	printList(head)

	// Modify node at position 3 to 99
	head = modifyAtPosition(head, 3, 99)

	fmt.Print("After Modifying Position 3 to 99: ")
	printList(head)
}

