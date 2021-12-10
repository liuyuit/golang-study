package main

import (
	"fmt"
	"time"
)

// 对于无缓冲 channel， 写入 channel 的 goroutine 和读取的 goroutine 形成了一种同步阻塞机制，
// 发送方需要确定数据被读取之后才能继续向下执行。读取方需要等到发送放写入数据才能继续向下执行
// 也就是双方都需要在 channel 中完成通信之后才能继续向下执行
func main() {
	c := make(chan int)

	go func(c chan int) {
		defer fmt.Println("new goroutine ended")
		intA := 1
		fmt.Println("new goroutine is running")
		fmt.Println("new goroutine ", intA)

		//time.Sleep(1 * time.Second)
		// 这里会阻塞，直到主 goroutine 读取了 channel 中的值
		c <- intA // 将值写入到 channel
		fmt.Println("new goroutine next")
	}(c)

	fmt.Println("main goroutine is running")

	var intB int
	// 从 c 中读取值，但并不接收
	//<-c
	// 这里会阻塞，直到 c 在协程中被赋值
	time.Sleep(1 * time.Second)
	intB = <-c
	fmt.Println("main goroutine ", intB)
	//time.Sleep(1 * time.Second)

	fmt.Println("main goroutine ended")
}
