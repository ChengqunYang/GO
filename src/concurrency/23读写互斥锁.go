package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	互斥锁是完全互斥的,但是也有很多的应用场景下是读多写少的,当我们并发的去读取一个资源的时候,是没有必要加锁的
这种场景下使用读写锁是更好的策略,读写锁在go语言中使用sync包中的RWMutex类型.
	读写锁分为两种:读锁和写锁,当一个goroutine获取到读锁之后,其他的goroutine如果是获取读锁会继续获得锁
如果是获取写锁就会等待,当一个goroutine获取写锁之后,其他的goroutine不管是获取读锁还是写锁都会等待.
*/

var(
	x = 0
	lock sync.Mutex
	wg sync.WaitGroup
	rwlock sync.RWMutex
	)

func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock()
	x++
	time.Sleep(time.Millisecond*5)
	//lock.Unlock()
	rwlock.Unlock()
}
func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwlock.RUnlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
