package main

import "fmt"

var (
	aa = "123"
	ss = 99
	cc = false
) //包内变量

func varibalefunc() {
	var s string = "hello"
	var a int = 55
	//fmt.Printf(a)
	fmt.Printf("%d %s\n", a, s)
}

func varibaleType() {
	a, b, c := 123, "hello", false
	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("hello")
	varibalefunc()
	varibaleType()
	fmt.Println(aa, ss, cc)
}
