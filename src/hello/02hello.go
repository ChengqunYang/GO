package main

import "fmt"

func main() {
	/*
		变量的声明
		1. 声明： var 变量名 类型  变量声明之后必须的使用
		2. 只是声明变量没有初始化，默认为0
		3. 在同一个{}里，变量的声明是唯一的
	*/
	var a int16
	a = 23 //变量的赋值
	println(a)
	var b, c int
	b, c = 20, 30
	//=后可以跟表达式
	var d int = 3 * 8 //声明变量的同时进行赋值
	println(d)
	fmt.Println(b, c)
	fmt.Println("hhhhh")
}
