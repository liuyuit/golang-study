package main

import "fmt"

// slice 是引用传递，传递的是当前数组的指针。切片本身就是存储的数组的引用。
func printArray(myArray []int) {
	for index, value := range myArray {
		fmt.Println(index, value)
	}

	myArray[0] = 111
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组
	fmt.Printf("%T\n", myArray)
	printArray(myArray)

	for index, value := range myArray {
		fmt.Println(index, value)
	}
}
