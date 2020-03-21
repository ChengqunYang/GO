package main

import (
	"fmt"
	"sync"
	"time"
)

//定义一个全局变量
var ch = make(chan int)
var wg sync.WaitGroup

//定义一个打印机,参数为字符串,按每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")

}
func Person1() {
	defer wg.Done()
	Printer("hello")
	ch <- 666 //向管道中写数据,因为该通道是一个无缓冲区的通道,
	// 所以如果没有别的线程从这个通道中拿走这个值,那么该线程将会被阻塞,后续的操作也会暂时不执行

	//Printer(" www ")
}
func Person2() {
	defer wg.Done()
	<-ch //从管道中拿数据,接收,如果通道没有数据它就会阻塞
	//因为只有当Person1打印完后才会向通道中写数据666.所以在此之前Person2都会被阻塞
	//当Person1执行完毕后,开始执行Person2
	Printer("world")
}
func main() {
	//新建两个协程,代表两个人,2个人同时使用打印机
	wg.Add(1)
	go Person1()
	wg.Add(1)
	go Person2()
	wg.Wait()
}
