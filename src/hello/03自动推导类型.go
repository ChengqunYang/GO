package main

import "fmt"

func main() {
	//自动推导类型
	var a = 1.534
	fmt.Println(a)

	fmt.Printf("%T\n", a)
	//自动推导类型的另一种方式
	b := 4.56
	fmt.Println("a = ", a, b)
	fmt.Println(b)
	c, d, e := 30, 23.5, 12
	fmt.Println(c, d, e)
}
