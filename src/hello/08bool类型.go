package main

import "fmt"

func main() {
	//1. 声明变量
	var a bool
	a = true
	fmt.Println(a)
	//2. 自动推导类型
	var b = false
	fmt.Println(b)

	c := true
	fmt.Println(c)
}
