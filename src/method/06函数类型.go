package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

//定义函数类型为hs
//函数也是一种数据类型,相当于将有两个int型参数,一个int型返回值的函数定义为hs这个类型
type hs func(int, int) int

func main() {
	//传统的调用方式
	var result int
	result = Add(1, 2)
	fmt.Println(result)
	//声明一个hs类型的函数变量
	var test hs
	//给这个test的变量赋值
	test = Add
	fmt.Println(test(10, 20))
	test = Minus
	fmt.Println(test(10, 4))

}
