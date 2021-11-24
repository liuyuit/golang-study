package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read book")
}

func (this *Book) WriteBook() {
	fmt.Println("write book")
}

func main() {
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	//w = b
	w.WriteBook()
}
