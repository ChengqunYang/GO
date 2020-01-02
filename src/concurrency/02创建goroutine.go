package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) //延时一秒
	}
}

//主协程退出了,子协程也就退出了
func main() {
	go newTask() //新建一个协程
	i := 0
	for {
		i++
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)

		if i == 5 {
			break
		}
	}

}
