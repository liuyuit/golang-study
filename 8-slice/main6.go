package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// [0, 2) 切片截取的下标是一个左闭右开的区间。省去左边表示从第 0 个开始，省去右边表示一直到最后一个
	s1 := s[0:2]
	fmt.Println(s)
	fmt.Println(s1)
	// 截取得到的切片和原来的切片指向同一个底层数组。
	s[0] = 100
	fmt.Println(s)  // [100 2 3]
	fmt.Println(s1) // [100 2]

	// 深拷贝
	s2 := make([]int, len(s))
	copy(s2, s)
	fmt.Println(s)  // [100 2 3]
	fmt.Println(s2) // [100 2 3]
}
