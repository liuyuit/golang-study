package main

import "fmt"

// 变量的构造 = pair
// type : static type int, string; concrete type interface 指向的具体类型，系统看得见的类型
// value
func main() {
	var a string
	// pair type:string  value:"abc"
	a = "abc"

	var allType interface{}
	// 这个赋值操作会把 a 变量的 pair 原样赋值给 allType。
	allType = a
	str, _ := allType.(string)
	fmt.Println(str)
}
