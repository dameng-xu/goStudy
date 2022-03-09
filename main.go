package main

import (
	"fmt"
	"github.com/dameng-xu/goStudy/function"
)

func main(){
	function.DeferFuncParameter()
	function.DeferFuncParameterV2()
	ret := function.DeferFuncReturn()
	fmt.Println(ret)
	ret2 := function.DeferFuncReturnV2()
	fmt.Println(ret2)
}
