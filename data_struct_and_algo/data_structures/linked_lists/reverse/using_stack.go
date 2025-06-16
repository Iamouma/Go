package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Function to reverse linked list using stack
func reverseUsingStack(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	// Stack to store node pointers
	stack := []*Node{}

	curr := head

	// Push all nodes to stack
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.next
	}

	// Pop the last node and make it the new head
	newHead := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	curr = newHead

	// Pop and link nodes in reverse order
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr.next = node
		curr = curr.next
	}

	// Set next of last node to nil
	curr.next = nil

	return newHead
}

// Utility function to print list
func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " ")
		head = head.next
	}
	fmt.Println()
}

// Main function to test
func main() {
	// Create list: 1 -> 2 -> 3 -> 4 -> 5
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

	// Reverse using stack
	reversedHead := reverseUsingStack(node1)

	fmt.Print("Reversed List: ")
	printList(reversedHead)
}

