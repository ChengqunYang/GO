package main

import (
	"fmt"
)

//当不需要再往channel中添加数据时,关闭channel,当管道关闭后,无法再发送数据
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

	//读取管道内容
	for true {
		//如果ok为true,说明管道并没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { //管道关闭
			break
		}
	}
}
