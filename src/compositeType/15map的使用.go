package main

import "fmt"

func main() {
	//map是无序的，我们无法确定他的返回顺序，所以每次打印结果都有可能不相同
	//定义一个map,键类型为int，值类型为string
	var m map[int]string
	fmt.Println("m = ", m)
	//对于map只有len没有cap
	fmt.Println("len=", len(m))

	//通过make创建
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	//通过make创建并指定长度(只是指定了可以有这么多)
	m3 := make(map[int]string, 10)
	m3[1] = "mike"
	m3[2] = "tony"
	m3[3] = "rose"
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))

	//初始化,键是唯一的
	m4 := map[int]string{
		1:"mike",2:"go",3:"c++"}
	fmt.Println(m4)
}
