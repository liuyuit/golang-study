package main

import "fmt"

func main() {
	// 定义一个 map, nil 的空指针
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("it is nil")
	}
	// 没有开辟空间之前，不能使用
	//myMap1["one1"] = "java" // panic: assignment to entry in nil map
	// 定义了之后，需要再开辟空间,开辟的空间如果用完，再次添加元素时，会自动开辟更多空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "java"
	myMap1["two"] = "python"
	myMap1["three"] = "c++"
	fmt.Println(myMap1)

	myMap2 := make(map[int]string, 10)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "c++"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one":   "php",
		"two":   "python",
		"three": "js",
	}
	fmt.Println(myMap3)
}
