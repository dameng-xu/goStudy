package main

import (
	"fmt"
	"github.com/dameng-xu/goStudy/dataStruct/linkList"
)

func main() {
	dll := linkList.DoublyLinkedList{}

	dll.InsertAtEnd(1)
	dll.InsertAtEnd(2)
	dll.InsertAtEnd(3)
	dll.InsertAtEnd(4)
	dll.InsertAtEnd(5)

	fmt.Println("Original list:")
	dll.Display()
	dll.Insert(4, 100)
	dll.PrintTrail()
	dll.Display()
}
