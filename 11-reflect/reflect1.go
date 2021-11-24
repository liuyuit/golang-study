package main

import (
	"fmt"
	"reflect"
)

func ReflectNum(arg interface{}) {
	fmt.Println(reflect.TypeOf(arg))
	fmt.Println(reflect.ValueOf(arg))
}

func main() {
	num := 3.14

	ReflectNum(num)
}
