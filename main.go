package main

import (
	"fmt"
	"github.com/dameng-xu/goStudy/dataStruct/linkList"
)

func main() {
	var l linkList.LinkList
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.Append(7)
	l.Add(0)
	l.Insert(5, 10000)
	l.Insert(0, 20000)
	l.PrintList()
	fmt.Println(l.Length())
}
