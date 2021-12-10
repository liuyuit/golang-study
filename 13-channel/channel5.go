package main

import (
	"fmt"
)

func fibonacii(c, quit chan int) {
	x, y := 1, 1

	for i := 0; i >= 0; i++ {
		// 在同一个 goroutine 中监控多个 channel。任何一个 channel 可读或可写都可以进入相应的 case 执行后续代码
		// 多路监控 channel
		select {
		case c <- x:
			// 如果 c 可写，则进入这个 case
			x = y
			y = x + y
			fmt.Println("x + y====================")

		case <-quit:
			// 如果 quit 中有数据可以读取到，则进入这个 case
			fmt.Println("quit")
			return
		default:
			// 如果上面的都没有触发，则进入 default
			fmt.Println("default")
		}
		fmt.Println("=========", i)
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 1
	}()

	//main go
	fibonacii(c, quit)
	fmt.Println("main is finished ---------------------------")
}
