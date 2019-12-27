package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	*Person //指针类型
	id      int
	addr    string
}

func main() {
	s1 := Student{
		Person: &Person{
			name: "tony",
			sex:  'm',
			age:  20,
		},
		id:   10,
		addr: "bj",
	}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)

	//先定义变量
	var s2 Student
	s2.Person = new(Person)
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 18
	s2.id = 110
	s2.addr = "xa"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)
}
