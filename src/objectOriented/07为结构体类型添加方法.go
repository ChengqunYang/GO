package main

import (
	"fmt"
)

//可以给任意自定义类型(包括内置类型,但不包括指针类型)添加相应的方法
type Person struct {
	name string
	sex  byte
	age  int
}

//只要接收者类型不一样,这个方法就算和别的方法同名,也是不同的方法,不会出现重复定义的错误
//带有接收者的函数叫做方法
func (p Person) PrintInfo() {
	fmt.Println("p = ", p)
}

//通过一个函数给成员赋值
//Person是接收者类型,它本身不能是指针类型,但是下面这种写法就可以
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}
func main() {
	//定义的同时初始化
	p := Person{"mike", 'm', 20}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person
	(&p2).SetInfo("yoyo", 'w', 26)
	p2.PrintInfo()
}
