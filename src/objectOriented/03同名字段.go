package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //匿名字段,那么默认Student就包含了Person的所有字段
	id     int
	addr   string
	name   string //和Person的name重名了
}

func main() {
	//定义一个变量
	var s Student
	//默认规则(就近原则,如果能在本作用域找到此成员,就操作此成员,如果没有找到,就找继承来的字段),操作的是Student的name
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Printf("s = %+v\n", s)

	//显示调用
	s.Person.name = "yoyo"
	fmt.Printf("s = %+v\n", s)
}
