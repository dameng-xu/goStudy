package main

import (
	"github.com/dameng-xu/goStudy/dataStruct/linkList"
)

func main() {
	l := linkList.LinkList{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(5)
	l.Add(6)
	l.Append(0)
	l.Append(100000000)
	l.Insert(0, 100)
	l.Insert(3, 101)
	l.Insert(9, 999)
	l.PrintList()
	l.Del(10)
	l.PrintList()
}
