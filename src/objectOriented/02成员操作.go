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
}

func main() {
	s1 := Student{
		Person: Person{"mike", 'm', 18},
		id:     1,
		addr:   "bj",
	}
	fmt.Printf("s1 = %+v\n", s1)
	s1.Person = Person{
		name: "tony",
		sex:  'w',
		age:  20,
	}
	fmt.Println(s1.name, s1.age, s1.id, s1.addr, s1.sex)
}
