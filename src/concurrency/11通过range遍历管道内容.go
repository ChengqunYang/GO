package main

import "fmt"

func main() {

	ch := make(chan int) //创建一个无缓冲的channel

	//创建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往channel中添加数据
		}
		//不需要再写数据时,关闭channel
		close(ch)
	}()
	//通过range的方式读取channel中的所有数据
	for num := range ch {
		fmt.Println("num= ", num)
	}
}
