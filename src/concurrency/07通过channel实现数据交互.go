package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个channel
	ch := make(chan string)
	defer fmt.Println("主协程结束")

	go func() {
		defer fmt.Println("子协程调用结束")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程i = ", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程,我工作完了"
	}()

	str := <-ch
	fmt.Println(str)
}
