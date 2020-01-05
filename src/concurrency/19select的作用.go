package main

import (
	"fmt"
)

/*
	go里面提供了一个关键字select,通过select可以监听channel的数据流动
	select的用法与switch语言非常的类似,由select开始一个新的选择块
	每个选择条件有case语句描述
	与switch语句可以选择任何可使用相等比较的条件相比,select有比较多的限制
	其中最大的一个就是每个case语句里面必须是一个io操作

	在一个select语句中,go语言会按顺序从头到尾评估每一个发送和接受的语句
	如果其中的任意一语句可以继续执行(即没有被阻塞),那么就从这些可以执行的
	语句中任意选择一条来执行

	如果没有任意一个语句可以继续执行(即所有的通道都被阻塞),那么有两种情况
	1. 如果给出了default语句,那么就会执行default语句,同时程序的执行也会从select
	语句后的语句中恢复
	2. 如果没有default语句,那么select语句将被阻塞,直到至少有一个通道可以进行下去
*/
//ch只写,quit只读
func fibo(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for true {
		//监听channel中数据的流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}
	}
}

func main() {
	//select实现斐波那契数列

	ch := make(chan int) //数字通信

	quit := make(chan bool) //程序是否结束

	//消费者,从channel读取内容
	go func() {

		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num = ", num)
		}
		//可以停止
		quit <- true
	}()
	//生产者,生产数字,写入channel
	fibo(ch, quit)

}
