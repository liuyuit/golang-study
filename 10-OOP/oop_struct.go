package main

import "fmt"

// 申明一种数据类型，是 int 的别名
type myint int

// 定义一种结构体，把多种复杂的数据类型放到一个结构中
type Book struct {
	title string
	auth  string
}

func main() {
	var a myint = 10
	fmt.Printf("%v : %T\n", a, a) // 10 : main.myint

	var book1 Book
	book1.title = "Golang"
	book1.auth = "Frank"
	fmt.Printf("%v \n", book1)
	changeBook(book1)
	fmt.Printf("%v \n", book1)

	changeBook2(&book1)
	fmt.Printf("%v \n", book1)
}

// 传递的是结构体的副本，不影响调用时传递的变量
func changeBook(book Book) {
	book.auth = "Zhangsang"
}

// 传递结构体的指针
func changeBook2(book *Book) {
	book.auth = "Zhangsang"
}
