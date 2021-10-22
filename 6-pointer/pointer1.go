package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pb = *pa
	*pa = temp
}

func main() {
	var a, b = 10, 20
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
