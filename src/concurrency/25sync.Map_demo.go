package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
	go语言中内置的map不是并发安全的,
存在多个goroutine之间的冲突
 会报错:fatal error: concurrent map writes
像这种场景下就需要加锁来保证并发的安全性了,go语言的sync 包中提供了一个开箱即用的并发安全版map
sync.Map 开箱即用就是说这种map不用像内置的map一样,使用make函数来初始化就能直接使用,同时sync.Map内置了
诸如:store\load\loadOrStore\Delete\Range等操作方法
*/

var m = make(map[string]int)
var lock sync.Mutex
var wg sync.WaitGroup
var m2 sync.Map

func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	m[key] = value
}

/*
func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key:= strconv.Itoa(n)
			lock.Lock()
			set(key,n)
			lock.Unlock()
			fmt.Printf("key =: %v,v =: %v\n",key,get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
*/
func main() {
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Printf("key =: %v,v =: %v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
