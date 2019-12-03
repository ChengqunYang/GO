package main

import "fmt"

//定义在函数外部的变量是全局变量,全局变量在任何地方都可以使用
var a int

func t() {
	fmt.Printf("test a = %d\n", a)
}
func main() {
	//局部变量和全局变量同名的话,就近原则使用变量
	a = 10
	fmt.Printf("a= %d\n", a)
	t()
}
