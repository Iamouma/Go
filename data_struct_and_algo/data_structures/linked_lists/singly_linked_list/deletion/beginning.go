package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// Function to delete the first node
func deleteAtBeginning(head *Node) *Node {
	// Step 1: Check if list is empty
	if head == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// Step 2: Store current head in temp
	temp := head

	// Step 3: Move head to next node
	head = head.next

	// Step 4: Clear temp (not strictly necessary in Go due to GC)
	temp.next = nil

	// Step 5: Return new head
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
	// Create linked list: 3 -> 12 -> 15 -> 18
	node1 := &Node{data: 3}
	node2 := &Node{data: 12}
	node3 := &Node{data: 15}
	node4 := &Node{data: 18}

	node1.next = node2
	node2.next = node3
	node3.next = node4

	head := node1

	fmt.Print("Original List: ")
	printList(head)

	// Delete first node
	head = deleteAtBeginning(head)

	fmt.Print("After Deletion: ")
	printList(head)
}

