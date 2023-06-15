package linkList

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkList struct {
	Head *Node
}

func (l *LinkList) Add(value interface{}) {
	l.Head = &Node{
		Value: value,
		Next:  l.Head,
	}
	return
}

func (l *LinkList) Append(value interface{}) {
	curNode := l.Head
	if curNode == nil {
		l.Add(value)
		return
	}
	// 如果向尾部添加，需要找到尾部不为null的节点，因为要使用node.next添加新的节点
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = &Node{
		Value: value,
		Next:  nil,
	}
	return
}

func (l *LinkList) Insert(index, value int) {
	if index == 0 {
		l.Add(value)
		return
	}

	curNode := l.Head
	curIdx := 1
	for curNode != nil {
		if curIdx == index {
			break
		}
		curNode = curNode.Next
		curIdx++
	}
	if curNode == nil {
		return
	}
	curNode.Next = &Node{
		Value: value,
		Next:  curNode.Next,
	}
	return
}

func (l *LinkList) Get(index int) *Node {
	curNode := l.Head
	if curNode == nil {
		return nil
	}
	curIdx := 0
	for curNode != nil {
		if curIdx == index {
			return curNode
		}
		curNode = curNode.Next
		curIdx++
	}
	return nil
}

func (l *LinkList) Del(index int) {
	preNode := l.Head
	if preNode == nil {
		return
	}
	if index == 0 {
		l.Head = preNode.Next
		return
	}

	curIdx := 1
	curNode := preNode.Next
	for curNode != nil {
		if curIdx == index {
			preNode.Next = curNode.Next
			return
		}
		preNode = curNode
		curNode = curNode.Next
		curIdx++
	}
	if curNode == nil {
		return
	}
}

func (l *LinkList) Reverse() {
	if l.Head == nil {
		return
	}

	pre := l.Head
	cur := pre.Next
	// 必须加，否则会形成循环链表
	pre.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	l.Head = pre
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
