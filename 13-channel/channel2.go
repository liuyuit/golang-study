package main

import (
	"fmt"
	"time"
)

// 无缓存 channel，发送者将数据写入 channel 中后可以继续向下执行，直到 channel 长度已满，此时会阻塞直到有空位写入数据
// 对于读取方，读取到数据后可以继续向下执行，如果 channel 已空，会阻塞直到有数据可以读取。
// 当 channel 已满，再向里面写入会阻塞
// 当 channel 为空，再从中读取会阻塞
func main() {
	c := make(chan int, 3)
	fmt.Printf("len(c) = %d, cap(c) = %d \n", len(c), cap(c))

	go func(c chan int) {
		defer fmt.Println("sub goroutine ended")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("sub goroutine i = %d, len(c) = %d, cap(c) = %d \n", i, len(c), cap(c))
		}
	}(c)

	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		//time.Sleep(1 * time.Second)
		fmt.Printf("main goroutine i = %d, len(c) = %d, cap(c) = %d \n", <-c, len(c), cap(c))
	}

	fmt.Println("main goroutine is ended")
}