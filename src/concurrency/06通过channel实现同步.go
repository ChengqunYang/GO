package main

import (
	"fmt"
	"time"
)

//定义一个全局变量
var ch = make(chan int)

//定义一个打印机,参数为字符串,按每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")

}
func Person1() {
	Printer("hello")
	ch <- 666 //向管道中写数据
}
func Person2() {

	<-ch //从管道中拿数据,接收,如果通道没有数据它就会阻塞
	//因为只有当Person1打印完后才会向通道中写数据666.所以在此之前Person2都会被阻塞
	//当Person1执行完毕后,开始执行Person2
	Printer("world")
}
func main() {
	//新建两个协程,代表两个人,2个人同时使用打印机
	go Person1()
	go Person2()
	for {
	}
}
