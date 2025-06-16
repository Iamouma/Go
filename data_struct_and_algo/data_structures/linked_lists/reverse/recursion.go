package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Recursive function to reverse the linked list
func reverseRecursive(head *Node) *Node {
	// Base case: empty list or last node
	if head == nil || head.next == nil {
		return head
	}

	// Recursively reverse the rest of the list
	newHead := reverseRecursive(head.next)

	// Make the next node point back to the current node
	head.next.next = head
	head.next = nil // set current node's next to nil

	return newHead
}

// Utility function to print the list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

// Main function to test
func main() {
	// Create Linked List: 1 -> 2 -> 3 -> 4 -> 5
	node1 := &Node{data: 1}
	node2 := &Node{data: 2}
	node3 := &Node{data: 3}
	node4 := &Node{data: 4}
	node5 := &Node{data: 5}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	fmt.Print("Original List: ")
	printList(node1)

	// Reverse using recursion
	reversedHead := reverseRecursive(node1)

	fmt.Print("Reversed List: ")
	printList(reversedHead)
}

