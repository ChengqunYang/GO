package main

import "fmt"

type Teacher struct {
	id int
	name string
	age int
}
func main() {
	//定义一个结构体普通变量
	var t Teacher
	//操作成员,需要使用.运算符
	t.name = "ma"
	t.age = 18
	t.id = 119
	fmt.Println(t)

	//指针有合法指向后,才操作成员
	//先定义一个普通结构体变量
	var tea Teacher
	//在定义一个指针变量保存tea的地址
	var p1 *Teacher
	p1 = &tea
	//通过指针操作成员 p1.id 和(*p1).id完全等价,只能使用.运算符
	p1.age = 28
	p1.name = "wang"
	(*p1).id = 110
	fmt.Println(*p1)

	//通过new来申请一个结构体
	p2 := new(Teacher)
	p2.id = 911
	p2.name = "li"
	p2.age = 45
	fmt.Println(*p2)

	//同类型的结构体的比较和赋值
	fmt.Println(p1==p2)
	p1 = p2
	fmt.Println(p1==p2)
}
