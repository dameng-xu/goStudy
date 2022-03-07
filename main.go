package main

import (
	"fmt"
	"github.com/dameng-xu/goStudy/dataStruct/dataSort"
)

func main(){
	s := []int{1, 2, 2, 1, 22, 888 , 11, 1272, 11, 227, 99, 287, 162, 67, 263}
	fmt.Println(s)
	dataSort.MergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}
