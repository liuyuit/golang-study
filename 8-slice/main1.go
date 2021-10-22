package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝

	/*	//for index, value := range myArray {
		//	fmt.Println("index = ", index, ", value = ", value)
		//}*/

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	myArray[0] = 111
}

func main() {
	//固定长度的数组
	var myArray1 = [4]int{}
	var myArray2 = [10]int{23, 12, 123234, 12}
	printArray(myArray1)
	//printArray(myArray2)

	//查看数组的数据类型
	fmt.Printf("%T\n", myArray2)

	var myArray3 = []int{1, 2, 3}

	fmt.Println(" ------ ")
	for index, value := range myArray3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
}
