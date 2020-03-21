package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
	代码中的加锁操作因为涉及到内核态的上下文切换会比较耗时,代价比较高,
针对基本数据类型,我们还可以使用原子操作来保证并发安全,因为原子操作是go语言提供的方法,
它在用户态就可以完成,因此性能比加锁要好,go语言中的原子操作作为内置的标准库sync.atomic提供
*/

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
/*	lock.Lock()
	x++
	lock.Unlock()*/
atomic.AddInt64(&x,1)
	wg.Done()
}
func main() {
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
