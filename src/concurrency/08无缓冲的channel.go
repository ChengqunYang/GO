package main

import (
	"fmt"
	"time"
)

func main() {
	//无缓存的channel,没有存储值的功能,每次放进去的数据只有在被拿走后才可以继续往进去放
	ch := make(chan int, 0)
	fmt.Printf("len = %d,cap = %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i=", i)
			ch <- i //将i写入channel中
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		/*num:=<-ch //管道中的数据取出来
		fmt.Println(num)*/
		fmt.Println(<-ch)
	}
}
