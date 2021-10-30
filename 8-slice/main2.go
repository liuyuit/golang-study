package main

import "fmt"

// 仍然是值拷贝
func printArray(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println(index, value)
	}

	myArray[0] = 11
}

func main() {
	var myArray1 [10]int

	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println(index, value)
	}

	fmt.Printf("%T\n", myArray1)
	fmt.Printf("%T\n", myArray2)
	fmt.Printf("%T\n", myArray3)

	printArray(myArray3)
	//printArray(myArray2) // cannot use myArray2 (type [10]int) as type [4]int in argument to printArray

	for index, value := range myArray3 {
		fmt.Println(index, value)
	}
}
