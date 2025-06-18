package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// DeleteNode deletes a node with the given key in a circular linked list
func DeleteNode(last *Node, key int) *Node {
	if last == nil {
		fmt.Println("List is empty.")
		return nil
	}

	// Only one node
	if last == last.next && last.data == key {
		return nil
	}

	var curr = last.next
	var prev = last

	// Traverse and search for the node to delete
	for curr != last {
		if curr.data == key {
			break
		}
		prev = curr
		curr = curr.next
	}

	// If the key is found
	if curr.data == key {
		prev.next = curr.next
		// If deleting the last node, update last
		if curr == last {
			last = prev
		}
		curr = nil // Optional: GC handles this
	} else if last.data == key {
		// If the last node is the one to be deleted
		prev.next = last.next
		last = prev

