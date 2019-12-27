package main

import "fmt"

//实现两个数相加
//面向过程
func add(a int, b int) int {
	return a + b
}

type long int

//面向对象,方法:给某个类型绑定一个函数
//tmp叫接收者,接受者就是传递的一个参数
func (tmp long) add1(other long) long {
	return tmp + other
}

func main() {
	var result int
	result = add(3, 4) //普通函数调用方式
	fmt.Println(result)

	//定义一个变量
	var a long = 2
	//调用方法格式: 变量名.函数(所需参数)
	r := a.add1(8) //通过一个接收者来调用该函数(面向对象就是换了一种表现形式)
	fmt.Println(r)
}
