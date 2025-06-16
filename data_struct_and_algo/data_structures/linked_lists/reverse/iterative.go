package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// function to reverse the linked list
func reverseList(head *Node) *Node {
	var prev *Node = nil
	curr := head

	for curr != nil {
		// store the next node
		next := curr.next

		// reverse the current node's pointer
		curr.next = prev

		// move pointers one position forward
		prev = curr
		curre = next
	}

	// prev will be the new head
	return prev
}

// utility function to print the list
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

	// Reverse the list
	reversedHead := reverseList(node1)

	fmt.Print("Reversed List: ")
	printList(reversedHead)
}

