package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"Name" doc:"my name"`
	Sex string `info:"gender"`
}

func findTag(str interface{}) {
	// 结构体的全部元素
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println(tagInfo, tagDoc)
	}
}

func main() {
	var rem resume

	findTag(&rem)
}
