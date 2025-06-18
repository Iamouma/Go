package main

import "fmt"

// Node defines a node of the doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// DeleteAtEnd deletes the last node and returns the head
func DeleteAtEnd(head *Node) *Node {
	if head == nil {
		return nil // List is empty
	}

	if head.next == nil {
		// Only one node
		head = nil
		return nil
	}

	// Traverse to the last node
	curr := head
	for curr.next != nil {
		curr = curr.next
	}

	// Remove last node
	curr.prev.next = nil
	curr = nil // Let Go's GC clean it up

	return head
}

// PrintList prints the doubly linked list from head to end
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.data)
		head = head.next
	}
	fmt.Println()
}

// main function to test deletion at end
func main() {
	// Create DLL: 1 <-> 2 <-> 3
	head := &Node{data: 1}
	second := &Node{data: 2}
	third := &Node{data: 3}

	head.next = second
	second.prev = head
	second.next = third
	third.prev = second

	fmt.Print("Original List: ")
	PrintList(head)

	// Delete last node
	head = DeleteAtEnd(head)

	fmt.Print("After Deletion at End: ")
	PrintList(head)
}

