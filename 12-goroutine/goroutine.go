package main

import (
	"fmt"
	"time"
)

// 子 goroutine
func newTask() {
/*	for i := 0; i < 10; i++ {  // 无输出
		fmt.Println(i)
	}*/

	i := 0
	for {
		i++
		fmt.Printf("new goroutine i = %d \n", i)
		time.Sleep(1 * time.Second)
	}
}

// main goroutine
func main() {
	go newTask()

	fmt.Println("main goroutine exited")

	// 如果主进程退出了，那么子进程也会退出。后面这段代码可以保证主进程不会退出
	i := 0
	for {
		i++
		fmt.Printf("main goroutine i=%d \n", i)
		time.Sleep(1 * time.Second)
	}
}
