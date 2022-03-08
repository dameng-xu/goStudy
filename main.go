package main

import (
	"fmt"
	"github.com/dameng-xu/goStudy/dataStruct/dataSort"
)

func main(){
	//var l linkList.SingleLinkList
	//l.Add(1)
	//l.Add(12)
	//l.Add(13)
	//l.Add(15)
	//l.Add(16)
	//l.Add(19)
	//l.Add(199)
	//l.Add(126)
	//l.Travel()
	//fmt.Println("****************")
	//l.Delete(12)
	//l.Delete(19)
	//l.Delete(126)
	//l.Travel()

	l := []int{
		1, 109, 18, 10, 28, 17, 272,18, 17, 277, 17, 277, 28, 19,111,11111,29292,0,-1,
	}
	dataSort.QuickSort(l, 0, len(l)-1)
	fmt.Println(l)
}
