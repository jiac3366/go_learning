package main

import "fmt"

func mapfunc() int {
	print("I am mapfunc")
	return 1
}

func main() {
	myMap := make(map[string]string, 10)
	fmt.Printf("map is %v\n", myMap)
	myMap["A"] = "A_value"

	fmt.Print(myMap["A"])
	value, exist := myMap["A"]
	if exist {
		println(value)
	}

	funcMap := make(map[string]func() int)
	funcMap["funckey"] = mapfunc
	// f := funcMap["funckey"]
	// 下面可以观察大写Print 和小写print差别 大写Print会首先输出 接着到print
	// fmt.Printf("%d", f()) //I am mapfunc
	// fmt.Print(f()) // I am mapfunc
	// fmt.Println(f()) //I am mapfunc
	// print(f())   //I am mapfunc
}
