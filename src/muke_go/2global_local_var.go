package main
var Global = "全局变量" //首字母大写 包外可访问
var local = "全局变量" //仅仅此包才可访问

// go可以对类型进行推断 数字会被理解为int float64 其他类型就要显示声明
var (
	first string = "abc"
	second int = 32
)

// 同作用域下，变量只能声明一次
var aa = "hello"
func func2()  {
	aa:= "45"
	println(aa)
}