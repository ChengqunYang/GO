package main

import (
	"fmt"
	"time"
)

//三种延时的方法
func main() {
	fmt.Println("开始")
	//延时两秒后,打印一句话
	/*timer := time.NewTimer(2 * time.Second)
	<-timer.C
	fmt.Println("时间到")*/

	/*	time.Sleep(2*time.Second)
		fmt.Println("时间到")*/
	<-time.After(2 * time.Second) //定时两秒,阻塞两秒,2秒后产生一个事件,往channel中写数据
	fmt.Println("时间到")
}
