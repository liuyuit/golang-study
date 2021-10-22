package main

import "fmt"

func returnFuc() int {
	fmt.Println("returnFuc called..")
	return 0
}

func deferFunc() int {
	fmt.Println("deferFunc called..")
	return 0
}

func main() {
	// returnFuc called..
	// deferFunc called..
	// return 语句会先执行，return 执行后才代表函数执行完了，然后才会执行 defer
	returnFuc()
	deferFunc()
}
