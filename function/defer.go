package function

import "fmt"

func DeferFuncParameter(){
	var aInt = 1
	defer fmt.Println(aInt)
	aInt = 2
	return
}

func printArray(arr *[3]int){
	for i := range arr{
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func DeferFuncParameterV2(){
	var aArray = [...]int{1, 2, 3}
	defer printArray(&aArray)

	aArray[0] = 10
	return
}

func DeferFuncReturn()(result int){
	i := 1

	defer func() {
		result ++
	}()

	return i
}

func DeferFuncReturnV2() int{
	var i int

	defer func() {
		i ++
	}()

	return i
}
