package main

import (
	"fmt"
)

func main() {
	//1. 字符,单引号   字符串,双引号
	//2. 字符,往往只有一个字符,转义字符除外'\n'  字符串又一个或者多个字符组成
	//3. 字符串都隐藏了一个结束符'\0'  如"a" 是由'a'和'\0'组成的
	var ch byte
	var str string
	ch = 'w'
	str = "什么嘛"
	println("ch=", ch, "str=", str)

	str = "hello world"
	//只打印其中的一个字符
	fmt.Printf("str[1]= %c\n", str[0])
	fmt.Println(len(str))
}
