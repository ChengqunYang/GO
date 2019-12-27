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
	var s1 Student = Student{
		Person: Person{"mike", 'm', 18},
		id:     12,
		addr:   "bj",
	}
	fmt.Println(s1)

	//自动推导类型
	s2 := Student{
		Person: Person{"tony", 'w', 19},
		id:     21,
		addr:   "xa",
	}
	fmt.Printf("s2 = %+v\n", s2)

	//制定成员初始化,没有初始化的成员默认为0
	s3 := Student{
		id: 1,
	}
	fmt.Printf("s3 = %+v\n", s3)

	s4 := Student{
		Person: Person{name: "yoyo"},
		id:     3}
	fmt.Printf("s4 = %+v\n", s4)
}
