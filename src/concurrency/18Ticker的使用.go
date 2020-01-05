package main

import (
	"fmt"
	"time"
)

/*
	ticker是一个定时触发的计时器,它会以一个间隔(interval)往channel中发送一个
时间(当前时间),而channel的接收者可以以固定的时间间隔从channel中读取事件
*/
func main() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for true {
		if i == 5 {
			ticker.Stop()
			break
		}
		i++
		fmt.Println("i=", i)
		<-ticker.C
	}
}
