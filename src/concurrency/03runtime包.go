package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccc")
	//return   //终止此函数 dddddddd将不会被打印
	runtime.Goexit() //终止了此协程
	fmt.Println("ddddddddddd")
}
func main() {
	//1.runtime.Gosched()用于让出cpu时间片,
	//让出当前goroutine的执行权限,调度器安排其他等待的任务运行,
	//并在下次某个时候从该位置恢复执行
	/*	go func() {
			for i:=0; i<5 ; i++ {
				fmt.Println("hello go")
			}
		}()
		for i := 0; i < 2; i++ {
			//让出时间片,让别的先执行
			runtime.Gosched()
			fmt.Println("hello")
		}*/
	//2.runtime.Goexit()用于结束当前协程
	/*	go func() {
			fmt.Println("aaaaaaaaaaaa")
			//调用了别的函数test
			test()
			fmt.Println("bbbbbbbbbbbb")
		}()
		for {

		}*/
	//runtime.GOMAXPROCS()用于设定cpu核数,返回值为当前电脑cpu的核数
	n := runtime.GOMAXPROCS(3) //设定当前任务以三核处理(三个cpu核心来轮换时间片来执行)
	fmt.Println(n)

}
