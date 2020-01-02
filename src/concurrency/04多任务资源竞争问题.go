package main

import (
	"fmt"
	"time"
)

//定义一个打印机,参数为字符串,按每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")

}
func main() {
	//新建两个协程,代表两个人,2个人同时使用打印机
	go Printer("hello")
	go Printer("world")
	//hweolrllod   执行结果表明这两个协程之间存在有冲突
	/*	go Printer("go 1111")
		//调用打印机打印字符
		Printer("hello")
		//特地不让主协程结束,死循环*/
	for {
	}
}
