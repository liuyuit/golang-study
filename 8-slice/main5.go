package main

import "fmt"

func main() {
	// 切片就是指向底层数组头元素的指针
	// 对于 slice 来说， len 表示当前切片中所有元素的个数，而 cap 是当前 slice 所指向的底层数组所能容纳的元素的个数。
	// 如果当前 cap 不够用了，再去 append 会自动动态扩充底层数组的容量
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d, cap=%d, %v \n", len(numbers), cap(numbers), numbers) // len=3, cap=5, [0 0 0]

	numbers = append(numbers, 1)
	fmt.Printf("len=%d, cap=%d, %v \n", len(numbers), cap(numbers), numbers) // len=4, cap=5, [0 0 0 1]

	numbers = append(numbers, 2)
	fmt.Printf("len=%d, cap=%d, %v \n", len(numbers), cap(numbers), numbers) // len=5, cap=5, [0 0 0 1 2]

	numbers = append(numbers, 3)
	fmt.Printf("len=%d, cap=%d, %v \n", len(numbers), cap(numbers), numbers) // len=6, cap=10, [0 0 0 1 2 3]

	numbers2 := make([]int, 3)
	fmt.Printf("numbers2 len=%d, cap=%d, %v \n", len(numbers2), cap(numbers2), numbers2) // len=6, cap=10, [0 0 0 1 2 3]
	numbers2 = append(numbers2, 1)
	fmt.Printf("numbers2 len=%d, cap=%d, %v \n", len(numbers2), cap(numbers2), numbers2) // len=6, cap=10, [0 0 0 1 2 3]
}
