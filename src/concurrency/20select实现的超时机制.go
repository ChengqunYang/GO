package main

import (
	"fmt"
	"time"
)

func main() {
	//有时候会出现goroutine阻塞的情况,那么我们如何避免整个程序进入阻塞的情况呢?
	//我们可以使用select来实现超时,通知如下的方式实现
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for true {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				fmt.Println("超时了")
				quit <- true
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit //防止程序执行完上面的for循环后直接结束,确保是因为超时了,才结束
	fmt.Println("程序结束")
}
