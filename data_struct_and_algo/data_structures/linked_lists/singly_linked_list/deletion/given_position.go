package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Function to delete a node at a given position (1-based index)
func deleteAtPosition(head *Node, position int) *Node {
	// If list is empty
	if head == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// If position is 1, delete the head node
	if position == 1 {
		temp := head
		head = head.next
		temp.next = nil
		return head
	}

	// Traverse to the node before the position
	current := head
	for i := 1; i < position-1 && current != nil; i++ {
		current = current.next
	}

	// If position is invalid (too large)
	if current == nil || current.next == nil {
		fmt.Println("Position out of bounds.")
		return head
	}

	// Skip the node at the given position
	temp := current.next
	current.next = temp.next
	temp.next = nil

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

	// Delete node at position 2
	head = deleteAtPosition(head, 2)

	fmt.Print("After Deleting at Position 2: ")
	printList(head)
}

