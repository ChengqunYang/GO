package main

import "fmt"

func MyFunc06() int {
	return 666
}

//go语言官方;推荐写法,给返回值一个名字
func MyFunc07() (result int) {
	result = 6666
	return
}

//定义有多个返回值的函数
func MyFunc08() (int, int, int) {
	return 1, 2, 3
}
func main() {
	//无参有返回值函数的调用
	var a int
	a = MyFunc06()
	fmt.Println(a)
	b := MyFunc06()
	fmt.Println(b)
	fmt.Println(MyFunc07())

	//调用有多个返回值的函数
	d, e, f := MyFunc08()
	fmt.Println(d, e, f)
}
