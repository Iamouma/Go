package main

import "fmt"

// Node strcuture
type Node struct {
	data int
	next *Node
}

// Recursive function to search for a key in the linked list
func searchRecursive(head *Node, key int) bool {
	//  Base case: if the head is nil, key is not found
	if head == nil {
		return false
	}

	// if current node's data matches the key
	if head.data == key {
		return true
	}

	// recurse on the next node
	return searchRecursive(head.next, key)
}

func main() {
	// Create linked list: 6 -> 21 -> 17 -> 30 -> 10 -> 8
	node1 := &Node{data: 6}
	node2 := &Node{data: 21}
	node3 := &Node{data: 17}
	node4 := &Node{data: 30}
	node5 := &Node{data: 10}
	node6 := &Node{data: 8}

	node1.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6

	key := 13

	if searchRecursive(node1, key) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
