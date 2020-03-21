package main

import (
	"fmt"
	"sync"
)

/*
	互斥锁是一种常用的控制共享资源访问的方法,它能够保证同一个时间只有一个goroutine可以访问
公共的资源.go语言中使用sync包中的Mutex类型来实现互斥锁.
*/
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	for i := 0; i < 5; i++ {
		go add()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(x)
}
