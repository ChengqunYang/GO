package main

import "fmt"

func main() {
	//1. 定义一个字符串
	var a string
	a = "Hello"
	b := "哈哈哈"
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)
	fmt.Println(len(b))
}
