package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 相对路径是相对执行命令的目录
	// tty pair <type *os.File> <value "11-reflect/tty" 文件描述符>
	tty, err := os.OpenFile("11-reflect/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error")
		fmt.Printf("error %v", err)
		return
	}

	// interface  r pair <type > <value>， pair 还未赋值
	var r io.Reader
	//  r pair <type *os.File> <value "11-reflect/tty" 文件描述符>
	r = tty

	// interface  w pair <type > <value>， pair 还未赋值
	var w io.Writer
	//  w pair <type *os.File> <value "11-reflect/tty" 文件描述符>
	w = r.(io.Writer)
	w.Write([]byte("this is a test!! \n"))
}
