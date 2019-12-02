package main

import (
	"fmt"
)

func main() {
	a := 10
	str := "mike"
	//匿名函数,没有函数名字
	//采用自动推到类型,简化书写,此时f1会根据后面的函数自定类型(无参无返回值)
	f1 := func() { //这是函数定义
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	//调用
	f1()

	//给一个函数类型起别名
	type FuncType func() //函数没有参数没有返回值
	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数并调用
	func() {
		fmt.Printf("a = %d,str = %s\n", a, str)
	}() //后面的()代表调用该函数

	//带参数的匿名函数
	func(i, j int) {
		fmt.Printf("i = %d,j = %d\n", i, j)
	}(2, 3)

	//匿名函数,有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i < j {
			max = j
			min = i
		} else {
			max = i
			min = j
		}
		return
	}(10, 20)
	fmt.Println(x, y)
}
