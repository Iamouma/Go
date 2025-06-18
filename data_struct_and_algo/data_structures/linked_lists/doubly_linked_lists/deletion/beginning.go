package main

import "fmt"

// Node defines a node of the doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// DeleteAtBeginning deletes the first node and returns the new head
func DeleteAtBeginning(head *Node) *Node {
	if head == nil {
		return nil // List is empty
	}

	newHead := head.next
	if newHead != nil {
		newHead.prev = nil
	}

	head = nil // Optional: Let Go's GC clean up
	return newHead
}

// PrintList prints the doubly linked list from head to end
func PrintList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.data)
		head = head.next
	}
	fmt.Println()
}

// main function to test deletion at beginning
func main() {
	// Create a simple DLL: 1 <-> 2 <-> 3
	head := &Node{data: 1}
	second := &Node{data: 2}
	third := &Node{data: 3}

	head.next = second
	second.prev = head
	second.next = third
	third.prev = second

	fmt.Print("Original List: ")
	PrintList(head)

	// Delete first node
	head = DeleteAtBeginning(head)

	fmt.Print("After Deletion at Beginning: ")
	PrintList(head)
}

