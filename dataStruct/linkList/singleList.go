package linkList

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkList struct {
	Head *Node
}

func (l *LinkList) Append(value int) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	if l.Head == nil {
		l.Head = &newNode
		return
	}
	lastNode := l.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = &newNode
	return
}

func (l *LinkList) Add(value int) {
	newNode := Node{
		Value: value,
		Next:  nil,
	}

	head := l.Head
	if head == nil {
		l.Head = &newNode
		return
	}
	newNode.Next = head
	l.Head = &newNode
	return
}

func (l *LinkList) Insert(index, value int) {
	if index < 0 || l.Length() < index {
		return
	}
	if index == 0 {
		l.Add(value)
		return
	}
	newNode := Node{
		Value: value,
		Next:  nil,
	}
	curNode := l.Head
	curIndex := 1
	for curIndex < index {
		curNode = curNode.Next
		curIndex++
	}
	newNode.Next = curNode.Next
	curNode.Next = &newNode
	return
}

func (l *LinkList) Length() int {
	var length int
	curNode := l.Head
	for curNode != nil {
		length++
		curNode = curNode.Next
	}
	return length
}

func (l *LinkList) PrintList() {
	currentNode := l.Head

	fmt.Print("LinkedList: ")
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Value)
		currentNode = currentNode.Next
	}
	fmt.Println("null")
}
