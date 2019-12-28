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
	//假如,结构体是一个指针变量,它能够调用那些方法,这些方法就是一个集合,称为方法集
	p := &Person{}
	p.SetInfoPointer()
	//内部做了转化,先把(*p)转化为p后再调用,等价于上面
	(*p).SetInfoPointer()

	//内部做了转化,先把指针p转化为*p后调用
	p.SetInfoValue()
	(*p).SetInfoValue()

	//用实例value和pointer调用方法,不受方法集的约束,编译器确总是查找全部的方法,并自动转化为receiver实参
}
