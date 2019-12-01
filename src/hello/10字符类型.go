package main

import "fmt"

func main() {
	//1.定义一个字符类型的变量
	var a byte
	a = 97
	fmt.Printf("%c\n", a)
	//%c以字符类型打印,%d以整型打印

	var b byte
	b = 'A'
	fmt.Printf("%d\n", b)
	//大小写相差32,小写的在后面
}
