package main

import "fmt"

func main() {
	//1. 类型别名就相当于给当前类型起一个别名
	type bigint int64
	var a bigint
	fmt.Printf("a%T\n", a)
}
