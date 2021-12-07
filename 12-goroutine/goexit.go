package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 退出这个 goroutine 子线程
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()

	go func(a int, b int) bool {
		fmt.Println(a, b)
		// 调用这个函数并不是一个阻塞操作，所以上层没法得到返回值。 可以通过 channel 来进行进程间的通信
		return true
	}(1, 2)

	for {
		time.Sleep(1 * time.Second)
	}
}
