package main

import "fmt"

//此通道只能写,不能读
func producer(out chan<- int) {
	for i := 0; i <= 10; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num = ", num)
	}
}
func main() {
	//创建一个channel
	ch := make(chan int)

	//生产者,生产一些数字,将数字写入channel
	//子协程,用来生产
	go producer(ch)
	//消费者,从channel中读取内容,打印
	//主协程用来消费
	consumer(ch)

}
