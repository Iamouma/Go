package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// SearchNode searches for a value in the circular linked list
func SearchNode(last *Node, key int) bool {
	if last == nil {
		fmt.Println("List is empty.")
		return false
	}

	curr := last.next // Start from head
	for {
		if curr.data == key {
			return true
		}
		curr = curr.next
		if curr == last.next {
			break
		}
	}
	return false
}

// Helper function to insert at end
func InsertAtEnd(last *Node, data int) *Node {
	newNode := &Node{data: data}
	if last == nil {
		newNode.next = newNode
		return newNode
	}
	newNode.next = last.next
	last.next = newNode
	return newNode
}

// Helper function to print circular linked list
func PrintList(last *Node) {
	if last == nil {
		fmt.Println("List is empty.")
		return
	}
	curr := last.next
	for {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
		if curr == last.next {
			break
		}
	}
	fmt.Println()
}

func main() {
	var last *Node
	last = InsertAtEnd(last, 10)
	last = InsertAtEnd(last, 20)
	last = InsertAtEnd(last, 30)
	last = InsertAtEnd(last, 40)

	fmt.Print("List: ")
	PrintList(last)

	key := 30
	if SearchNode(last, key) {
		fmt.Printf("Element %d found.\n", key)
	} else {
		fmt.Printf("Element %d not found.\n", key)
	}

	key = 99
	if SearchNode(last, key) {
		fmt.Printf("Element %d found.\n", key)
	} else {
		fmt.Printf("Element %d not found.\n", key)
	}
}

