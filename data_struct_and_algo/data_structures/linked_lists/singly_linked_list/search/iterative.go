package main

import "fmt"

// Node structure
type Node struct {
	data int
	next *Node
}

// iterative function to search for a key in a linked list
func search(head *Node, key int) bool {
	current := head
	
	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}

	return false
}

func main() {
	// Create Linked List: 14 -> 21 -> 11 -> 30 -> 10
	node1 := &Node{data: 14}
	node2 := &Node{data: 21}
	node3 := &Node{data: 11}
	node4 := &Node{data: 30}
	node5 := &Node{data: 10}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5

	key := 14

	if search(node1, key) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
