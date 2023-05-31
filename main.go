package main

import (
	"github.com/dameng-xu/goStudy/dataStruct/heap"
)

func main() {
	var h heap.Heap
	h.Pop()
	h.List()
	h.Push(101)
	h.Push(99)
	h.Push(11)
	h.Push(1)
	h.Push(92)
	h.Push(1000)
	h.Push(12781)
	h.Push(28)
	h.List()
	h.Push(8202)
	h.List()
	h.Pop()
	h.List()
	h.Pop()
	h.List()
	h.Push(8202)
	h.List()
	h.Push(8202)
	h.List()
	h.Push(8202)
	h.List()
	h.Pop()
	h.List()
}
