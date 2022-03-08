package linkList

import "fmt"

type SingleNode struct {
	Data	int
	Next 	*SingleNode
}

type SingleLinkList struct {
	Head	*SingleNode
}

func (s *SingleLinkList) Add(d int)  {
	node := SingleNode{Data: d}
	head := s.Head
	if head == nil{
		s.Head = &node
	}else {
		for head.Next != nil{
			head = head.Next
		}
		head.Next = &node
	}
}

func (s *SingleLinkList) Travel()  {
	head := s.Head
	for head != nil{
		fmt.Println(head.Data)
		head = head.Next
	}
}

func (s *SingleLinkList) Reverse()  {
	var pre *SingleNode
	var next *SingleNode

	head := s.Head
	for head != nil{
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	s.Head = pre
}

func (s *SingleLinkList) Delete(d int)  {
	head := s.Head
	for head != nil && head.Data == d{
		head = head.Next
	}
	s.Head = head

	cur := head
	next := head
	for next != nil{
		if next.Data == d{
			cur.Next = next.Next
		}else{
			cur = next
		}
		next = next.Next
	}
}
