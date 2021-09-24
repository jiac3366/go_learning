package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fullString := "hello world"
	for index, value := range fullString {
		fmt.Println(index, string(value)) //使用stirng可以把ascii码转换成字符串
	}
}
