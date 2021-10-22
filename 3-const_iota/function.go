package main

import "fmt"

const (
	a = iota
	b = iota
)

const c = iota
const d = iota

const (
	e = iota * 2
	f = iota * 2
)

func main() {
	// a = 10
	fmt.Println(a, b, c, d, e, f)
}
