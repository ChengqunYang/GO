package main

import "fmt"

func main() {
	var a int
	a = 20
	fmt.Println(a)

	var b = 15
	fmt.Println(b)

	c := 10
	fmt.Println(c)

	d, e := 30, 40
	fmt.Println(d, e)

	//不同类型变量的定义
	var f int = 1
	var g float64 = 2.0
	fmt.Println(f, g)

	//常量
	const x int = 10
	fmt.Println("x = ", x)

	//常量的自动推导类型
	const y = 20.00
	fmt.Printf("y type is %T\n", y)
	fmt.Println("y = ", y)

}
