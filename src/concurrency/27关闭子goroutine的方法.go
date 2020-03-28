package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
为什么要使用context?
*/

var wg sync.WaitGroup

//var notify bool  //零值为false
var quit = make(chan bool, 1)

func f() {
	defer wg.Done()
	//无限循环
LOOP:
	for {
		fmt.Println("天空")
		time.Sleep(time.Millisecond * 500)
		/*if notify{
			//如果notify为true的话,退出循环
			break
		}*/
		select {
		case <-quit:
			break LOOP
		default:
			continue
		}
	}
}
func f2(cancel context.Context) {
	defer wg.Done()
	LOOP:
	//无限循环
	for {
		fmt.Println("天空")
		time.Sleep(time.Millisecond * 500)
		select {
		case<-cancel.Done():
			break LOOP
		default:

		}
	}
}

func main() {
	wg.Add(1)
	//go f()
	//time.Sleep(time.Second * 3)
	//如何通知子goroutine退出?

	/*方法一:借助全局变量
	notify = true
	wg.Wait()
	全局变量方式存在的问题：
	1. 使用全局变量在跨包调用时不容易统一
	2. 如果worker中再启动goroutine，就不太好控制了。
	*/

	/*
		方法二:借助channel
		quit <- true
		wg.Wait()
		1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel
	*/

	/*
		方法三:借助context
	*/
	cancel, cancelFunc := context.WithCancel(context.Background()) //造一个取消函数,参数为一个父级的context
	go f2(cancel)
	time.Sleep(time.Second * 3)
	cancelFunc()
	wg.Wait()

}
