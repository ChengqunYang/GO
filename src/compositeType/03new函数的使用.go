package main

import "fmt"

func main() {
	//a := 10
	var p *int
	//让P指向一个合法内存
	//p = &a

	//p是*int
	p = new(int)

	*p = 666
	fmt.Println("*p = ", *p)

	//自动推导类型
	q := new(int)
	*q = 777
	fmt.Println("q = ", *q)
}
