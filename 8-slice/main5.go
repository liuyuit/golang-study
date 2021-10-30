package main

import "fmt"

func main() {
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
