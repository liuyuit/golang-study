package main

import "fmt"

func main() {
	c := make(chan int)
	go func(c chan int) {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// 关闭一个 channel, 无法向已关闭的 channel 写入数据
		close(c)
	}(c)

	for {
		// ok 为 true 表示 channel 尚未关闭，否则已经关闭
		// 对于有缓存 channel 只有无数据且已经关闭才会 返回 false
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main goroutine is finished")
}
