package main

import "fmt"

/*func swap (a *int, b *int) (int, int) {
	var temp *int

	temp = a
	a = b
	b = temp
	return *a, *b
	//var a int = 10
	//var b int = 20
	//
	//a, b = swap(&a, &b)
}*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa

	//fmt.Println(*pb, pb, &pb) // temp 20 0xc0000100c0 0xc000006028
	*pa = *pb
	*pb = temp
}

func test(str *string) {
	fmt.Println(str, *str)
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int
	p = &a
	fmt.Println(&a, p, *p)

	var pp **int //二级指针
	//fmt.Println(pp) // <nil>

	pp = &p
	fmt.Println(pp, *pp, **pp)

	var str string = "hello"
	var pstr *string
	pstr = &str

	test(pstr)

	fmt.Println(str, pstr, *pstr)
}
