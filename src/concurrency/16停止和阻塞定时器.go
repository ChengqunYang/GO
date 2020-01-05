package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("定时器的时间到,子协程可以打印了")
	}()

	timer.Stop() //停止定时器
	for true {
	}
}
