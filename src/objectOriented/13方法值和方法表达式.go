package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer: %p, %v\n", p, *p)
}

func main() {
	p := Person{
		name: "mike",
		sex:  'm',
		age:  20,
	}
	fmt.Printf("main: %p,%v\n", &p, p)
	//传统的调用方式
	p.SetInfoPointer()
	p.SetInfoValue()
	//保存函数入口地址
	pFunc := p.SetInfoPointer //这个就是方法值,调用函数时无需再传递接受者
	pFunc()                   //等价于 p.SetInfoPointer()
	//方法值的好处是隐藏了接收者
	vFunc := p.SetInfoValue
	vFunc() //等价于p.SetInfoValue()

	//方法表达式
	f := (*Person).SetInfoPointer //因为此处并没有指出到底谁是接收者,仅仅提供了接收者的类型
	//所以在调用的时候仍需要将接收者传递过去
	f(&p)
}
