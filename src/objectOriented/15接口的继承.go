package main

import "fmt"

type Humaner interface {
	sayhi()
}
type Personer interface {
	Humaner //匿名字段,继承了Humaner接口,也就继承了Humaner接口中的函数
	sing(irc string)
}

type Student struct {
	name string
	id   int
}

//Student实现了sayhi函数
func (tmp *Student) sayhi() {
	fmt.Printf("Stundet [%s,%d] sayhi\n", tmp.name, tmp.id)
}
func (tmp *Student) sing(irc string) {
	fmt.Println("Student在唱着:", irc)
}
func main() {
	//定义一个接口类型的变量
	var i Personer
	s := &Student{
		name: "mike",
		id:   20,
	}
	i = s
	i.sayhi() //继承的Humaner接口中的方法
	i.sing("太阳从东方升起")
}
