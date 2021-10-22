package main

import "fmt"

func main() {
	//写入defer关键字, 会在 defer 所在函数结束前触发
	// defer 语句会压入栈区，先入栈的会后执行
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
