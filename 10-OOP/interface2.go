package main

import "fmt"

// 万能类型，空接口。 因为所有的基本类型 int float 32 struct 都实现了 interface{}
func MyFunc(aug interface{}) {
	fmt.Println("MyFunc is called")

	// interface 类型断言, 必须是 空接口才能用，
	// 判断是否为 string
	value, ok := aug.(string)
	if ok {
		fmt.Println("it's string")
		fmt.Printf("value type is %T == %v \n", value, value)
	} else {
		fmt.Println("it's not a string")
	}

	fmt.Println("========")
}

type Book struct {
	auth string
}

func main() {
	book := Book{auth: "Frank"}
	MyFunc(book)
	MyFunc(12)
	MyFunc("abc")
}
