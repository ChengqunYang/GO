package main

import (
	"fmt"
	"time"
)

/*
	timer是一个定时器,代表未来的一个单一事件,你可以告诉timer你要等待多长时间,
它提供一个channel,在将来的那个时间给timer.C这个channel提供了一个时间值
	验证timer.NewTimer(),时间到了,只会响应一次
*/
func main() {
	//创建一个定时器,设置时间为2秒,2秒后,往timer通道里面写内容
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间", time.Now())
	//两秒后,数据被写入channel中,再从通道中读取数据
	t := <-timer.C //channel没有数据前会阻塞
	fmt.Println(t)

}
