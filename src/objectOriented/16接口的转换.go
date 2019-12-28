package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}
type Personer interface { //超集
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

	var per Personer
	var hum Humaner
	per = &Student{
		name: "mike",
		id:   666,
	}
	per.sing("我的未来不是梦")
	//per = hum  //不可以
	hum = per //超集可以转子集
	hum.sayhi()
}
