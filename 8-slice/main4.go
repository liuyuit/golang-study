package main

import "fmt"

func main() {
	//slice1 := []int{1,2,3} // 3  [1 2 3]
	//var slice1 []int
	//slice1 = make([]int, 3) // 开辟空间

	//var slice1 []int = make([]int , 3)
	slice1 := make([]int, 3)
	fmt.Printf("%d  %v", len(slice1), slice1)
	slice1[0] = 1

	//判断一个silce是否为0
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
