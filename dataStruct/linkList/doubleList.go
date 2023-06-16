package linkList

import (
	"fmt"
)

// DNode represents a DNode in the doubly linked list
type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head *DNode
	tail *DNode
}

// InsertAtEnd inserts a new DNode with the given data at the end of the list
func (dll *DoublyLinkedList) InsertAtEnd(data interface{}) {
	newDNode := &DNode{data: data, prev: nil, next: nil}

	if dll.head == nil {
		dll.head = newDNode
		dll.tail = newDNode
		return
	}

	current := dll.head
	for current.next != nil {
		current = current.next
	}

	newDNode.prev = current
	current.next = newDNode
	dll.tail = newDNode
}

// Add insert a new DNode in the head place
func (dll *DoublyLinkedList) Add(data interface{}) {
	if dll.head == nil {
		dll.InsertAtEnd(data)
		return
	}
	dll.head = &DNode{
		prev: nil,
		data: data,
		next: dll.head,
	}
	return
}

func (dll *DoublyLinkedList) Insert(index int, data interface{}) {
	if dll.head == nil {
		return
	}
	if index == 0 {
		dll.Add(data)
		return
	}
	pre := dll.head
	curIdx := 1
	cur := pre.next

	for cur != nil {
		if curIdx == index {
			newNode := DNode{
				data: data,
				prev: pre,
				next: cur,
			}
			pre.next = &newNode
			cur.prev = &newNode
			return
		}
		pre = cur
		cur = cur.next
		curIdx++
	}
}

// Delete deletes the first occurrence of the DNode with the given data from the list
func (dll *DoublyLinkedList) Delete(index int) {
	pre := dll.head
	if pre == nil {
		return
	}
	if index == 0 {
		next := dll.head.next
		dll.head = next
		if next == nil {
			dll.tail = nil
			return
		}
		next.prev = nil
		return
	}
	curIndex := 1
	current := pre.next
	for current != nil {
		if curIndex == index {
			next := current.next
			pre.next = next
			if next != nil {
				next.prev = pre
				return
			}
			dll.tail = pre
			return
		}
		pre = current
		current = current.next
		curIndex++
	}
}

// Get there value in the index index
func (dll *DoublyLinkedList) Get(index int) interface{} {
	curNode := dll.head
	curIndex := 0
	for curNode != nil {
		if curIndex == index {
			return curNode.data
		}
		curNode = curNode.next
		curIndex++
	}
	return nil
}

// Reverse reverses the order of the DNodes in the list
func (dll *DoublyLinkedList) Reverse() {
	current := dll.head
	var prev *DNode

	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}

	dll.head = prev
}

// Display prints the elements of the list
func (dll *DoublyLinkedList) Display() {
	current := dll.head

	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}

	fmt.Println()
}

func (dll *DoublyLinkedList) PrintTrail() {
	fmt.Println("tail: ", dll.tail)
}
