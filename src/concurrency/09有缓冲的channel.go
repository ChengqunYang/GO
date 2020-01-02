package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个有缓存的channel,可以存储一定的数据,在容量未满之前,往里面添加数据并不会阻塞
	ch := make(chan int, 3)
	fmt.Printf("len = %d,cap = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i=", i)
			ch <- i //将i写入channel中
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //管道中的数据取出来
		fmt.Println(num)
	}
}
