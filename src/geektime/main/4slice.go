package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 15, 40}
	mySlice1 := make([]int, 0, 20) //推荐创建法
	mySlice2 := append(mySlice1, 2)
	fmt.Printf("myslice2 is %+v\n", mySlice2)

	// 下面无法修改value的值 因为go是值传递
	for _, value := range mySlice {
		// 临时的value
		value *= 2
	}
	fmt.Printf("my slice %+v", mySlice)

	// for index, _ := range mySlice { 等同-->
	for index := range mySlice {
		// 通过下标
		mySlice[index] *= 2
	}
	fmt.Printf("my slice %v", mySlice)
	fmt.Printf("my slice %+v", mySlice)
	// fmt.Sprintf
	// fmt.Errorf
}
