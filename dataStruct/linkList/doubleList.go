package linkList

type DoubleNode struct {
	Data	int
	Next 	*DoubleNode
	Pre 	*DoubleNode
}

type DoubleLinkList struct {
	Head 	*DoubleNode
}

func (d *DoubleLinkList) Add(data int)  {
	node := DoubleNode{Data: data}
	head := d.Head
	if head == nil{
		d.Head = &node
	}else {
		for head.Next != nil{
			head = head.Next
		}
		head.Next = &node
		node.Pre = head
	}
}
