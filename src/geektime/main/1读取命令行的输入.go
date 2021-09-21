package main

// go的package一定要是main
//go package不是按照目录顺序的 go文件所在目录名可以不是main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 输出命令行参数
	fmt.Println(os.Args)
	fmt.Println("the len of args is", len(os.Args))
	if len(os.Args) > 1 {
		fmt.Println("the os args is", os.Args)
	}

	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	//main()无返回值
	// os.Exit(-1)

	//name的默认值是world 如果命令行显示传 --name xxx xxx就会替代world
	name := flag.String("name", "world", "specify the name u wanna say")
	flag.Parse()
	fmt.Println("input args is", *name) // input args is jiac

}
