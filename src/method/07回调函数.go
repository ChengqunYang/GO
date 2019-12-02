package main

import "fmt"

type FuncType func(int, int) int

//实现加法
func add(a, b int) int {
	return a + b
}

//实现减法
func minus(a, b int) int {
	return a - b
}

//回调函数:函数有一个参数是函数类型,这个函数就是回调函数
//两个好处,1.实现多态2.先定义,后实现
//计算器,可以进行四则运算
func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	return
}
func main() {
	//用于实现多态,调用同一个接口,不同的表现
	a := Calc(1, 1, add)
	fmt.Println(a)

	b := Calc(5, 1, minus)
	fmt.Println(b)
}
