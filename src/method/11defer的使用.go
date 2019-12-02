package main

import "fmt"

func ff(x int) int {
	return 10 / x
}
func main() {

	//defer:延迟调用,main函数结束前一瞬间调用
	defer fmt.Println("bbbbbb")

	fmt.Println("aaaaaa")

	//多个defer执行顺序
	//如果一个函数中有多个defer语句,他们会以先进后出的顺序执行,哪怕函数或某个延迟调用发生错误,这些调用依旧会被执行
	defer fmt.Println("ddddddd")
	defer ff(0)
	defer fmt.Println("eeeeeee")
}
