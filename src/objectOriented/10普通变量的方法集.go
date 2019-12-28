package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

func (p Person) SetInfoValue() {
	p.name = "mike"
	p.sex = 'm'
	p.age = 18
	fmt.Println(p)
}

func (p *Person) SetInfoPointer() {
	p.name = "tony"
	p.sex = 'w'
	p.age = 20
	fmt.Println(*p)
}
func main() {
	//用实例value和pointer调用方法,不受方法集的约束,编译器确总是查找全部的方法,并自动转化为receiver实参
	p := Person{}
	//内部先把p转化为&p再调用,(&p).SetInfoPointer()
	p.SetInfoPointer()
	p.SetInfoValue()
	(&p).SetInfoPointer()
	(&p).SetInfoValue()
}
