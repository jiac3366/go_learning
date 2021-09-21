package main

import (
	"awesomeProject1/ch1"
) //导入包
import "unicode/utf8"
// rune 本质是 int32，一个 rune 四个字节 替代其他语言的char
// byte本质为unint8
func main() {
	println("hello") //string 只能和string 拼接
	println(len(`你好我是`))  //12 len()获取字节长度：和编码无关
	println(utf8.RuneCountInString("你好WPS"))  //5 用编码库来计算字符数量：和编码有关，
	println("func2")
	//-----2.go
	func2()
	println(Global)
	//println(lobal) //无法访问
	//-----ch1/ch1.go
	Func()
	ch1.Ch1()
	println(add(5,6))
	println(add2(4,5))  //fmt.Println大写： 导出输出
	slice()

}