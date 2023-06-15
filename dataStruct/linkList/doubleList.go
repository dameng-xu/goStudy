package linkList

import "fmt"

type DoubleNode struct {
	Data interface{}
	Pre  *DoubleNode
	Next *DoubleNode
}

type DoubleLinkList struct {
	Head *DoubleNode
	Tail *DoubleNode
}

func (d *DoubleLinkList) Append(value interface{}) {
	newNode := DoubleNode{
		Data: value,
		Pre:  nil,
		Next: nil,
	}
	if d.Head == nil {
		d.Head = &newNode
		d.Tail = &newNode
		return
	}
	d.Tail.Next = &newNode
	newNode.Pre = d.Tail
	d.Tail = &newNode
	return
}

func (d *DoubleLinkList) Add(value interface{}) {
	newNode := DoubleNode{
		Data: value,
		Pre:  nil,
		Next: nil,
	}
	if d.Head == nil {
		d.Head = &newNode
		d.Tail = &newNode
		return
	}

	newNode.Next = d.Head
	d.Head.Pre = &newNode
	d.Head = &newNode
	return
}

func (d *DoubleLinkList) PrintForward() {
	currentNode := d.Head

	fmt.Print("DoublyLinkedList (Forward): ")
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println()
}

func (d *DoubleLinkList) PrintBackward() {
	currentNode := d.Tail

	fmt.Print("DoublyLinkedList (Backward): ")
	for currentNode != nil {
		fmt.Printf("%d <- ", currentNode.Data)
		currentNode = currentNode.Pre
	}
	fmt.Println()
}
