package main

import "fmt"

func testa() {
	fmt.Println("aaaaaa")
}
func testb(x int) {
	//设置recover
	defer func() {
		//recover()
		//fmt.Println(recover())//可以打印错误信息
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //()的意思是调用此匿名函数

	var a [10]int
	a[x] = 111 //当x大于9时,导致数组越界,产生一个panic,导致程序崩溃
	fmt.Println("未越界,赋值成功")

}
func testc() {
	fmt.Println("ccccccc")
}
func main() {
	testa()
	testb(10)
	testc()

}
