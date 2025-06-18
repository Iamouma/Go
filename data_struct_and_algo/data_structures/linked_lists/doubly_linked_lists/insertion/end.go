package main

import "fmt"

// Node represents a node in a doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// InsertAtEnd inserts a new node at the end of the doubly linked list
func InsertAtEnd(head *Node, data int) *Node {
	newNode := &Node{data: data}

	// If the list is empty, new node becomes the head
	if head == nil {
		return newNode
	}

	curr := head
	// Traverse to the last node
	for curr.next != nil {
		curr = curr.next
	}

	// Set the next of last node to newNode and prev of newNode to last node
	curr.next = newNode
	newNode.prev = curr

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

func main() {
	var head *Node

	head = InsertAtEnd(head, 10)
	head = InsertAtEnd(head, 20)
	head = InsertAtEnd(head, 30)
	head = InsertAtEnd(head, 40)

	fmt.Print("Doubly Linked List: ")
	PrintList(head)
}

