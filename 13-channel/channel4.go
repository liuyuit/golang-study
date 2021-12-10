package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(c chan int) {
		for i := 0; i < 5; i++ {
			//time.Sleep(1 * time.Second)
			c <- i
		}
		close(c)
	}(c)

	for data := range c {
		fmt.Println(data)
	}
}
