package main

import (
	"fmt"
	"sync"
)

/*
	在编程的很多场景下,我们需要保证某些操作在高并发的场景下只执行一次,
	例如只加载一次配置文件,只关闭一次通道等
go语言中的sync包中提供了一个针对只执行一次场景的解决方案sync.once
*/

//关闭channel的练习
var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch1<-i
	}
	close(ch1)
}

func f2(ch1<-chan int,ch2 chan<-int)  {
	defer wg.Done()
	for true {
		x,ok:= <-ch1
		if !ok{
			break
		}
		ch2<-x*x
	}
	once.Do(func() {
		close(ch2)
	})
	//保证关闭通道的操作只执行一次
}
func main() {
	a:= make(chan int,100)
	b:= make(chan int ,100)
	wg.Add(3)
	go f1(a)
	//一个往通道里面写值,两个从通道里面取值并计算平方
	go f2(a,b)
	go f2(a,b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)

	}

}
