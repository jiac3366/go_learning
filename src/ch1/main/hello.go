package main

// go的package一定要是main
//go package不是按照目录顺序的 go文件所在目录名可以不是main

import (
	"fmt"
	"os"
)



func main() {
	// 输出命令行参数
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))
	if len(os.Args) > 1 {
		fmt.Println("fuck vs code", os.Args[1])
	}

	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	//main()无返回值
	// os.Exit(-1)
}
