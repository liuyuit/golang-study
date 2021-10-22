package main

import "fmt"

func main() {
	// defer 会往栈区插入一条数据，先入后出
	defer fmt.Println("end 1")
	defer fmt.Println("end 2")

	fmt.Println("go 1")
	fmt.Println("go 2")
}
