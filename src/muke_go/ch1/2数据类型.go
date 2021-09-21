package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//go语言没有char类型 只有rune类型
//复数 complex64实部和虚部是float32
//complex128 实部和虚部是float64
func euler() {
	e := 3 + 4i
	aa := cmplx.Abs(e)
	fmt.Println(aa)

	//以下2种写法等价
	fmt.Println(
		cmplx.Exp(1i*math.Pi)+1,         //欧拉公式：e的(i*Pi)次方 + 1 =0
		cmplx.Pow(math.E, 1i*math.Pi)+1, //输出(0+1.2246467991473515e-16i) 不为零 那么如何输出0呢？
	)

	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1) //输出(0.000+0.000i)
	fmt.Println("")
}

func force_switch_type_of_data() {
	//计算勾股定理
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //注意这里2次强制转换
	fmt.Println(c)
}

//常量类型
func const_type() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4 //const数值可以作为各种类型使用 因为sqrt()内部需要float64
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//
func enumerate() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	const (
		b  = 1 <<(10*iota) //iota = 0, 1, 2, 3 ...
		kb
		mb
		g
		tg
	)
	fmt.Println(b, kb, mb, g, tg)
}



func main() {
	euler()
	force_switch_type_of_data()
	const_type()
	enumerate()

}
