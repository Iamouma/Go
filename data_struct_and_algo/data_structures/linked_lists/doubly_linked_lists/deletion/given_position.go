package main

import "fmt"

// Node defines a node of the doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// DeleteAtPosition deletes a node at the given position (1-based)
func DeleteAtPosition(head *Node, position int) *Node {
	if head == nil || position <= 0 {
		return head // Invalid position or empty list
	}

	if position == 1 {
		return DeleteAtBeginning(head)
	}

	curr := head
	for i := 1; curr != nil && i < position; i++ {
		curr = curr.next
	}

	if curr == nil {
		return head // Position out of range
	}

	// Disconnect node
	if curr.prev != nil {
		curr.prev.next = curr.next
	}
	if curr.next != nil {
		curr.next.prev = curr.prev
	}
	curr = nil // Let Go's GC clean it

	return head
}

// DeleteAtBeginning is reused from earlier code
func DeleteAtBeginning(head *Node) *Node {
	if head == nil {
		return nil
	}

	newHead := head.next
	if newHead != nil {
		newHead.prev = nil
	}
	head = nil
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

// main function to test deletion at specific position
func main() {
	// Create DLL: 1 <-> 2 <-> 3 <-> 4
	head := &Node{data: 1}
	second := &Node{data: 2}
	third := &Node{data: 3}
	fourth := &Node{d

