package main

import "fmt"

type mystr string

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //结构体匿名字段,那么默认Student就包含了Person的所有字段
	int    //基础类型的匿名字段
	mystr  //自定义类型的匿名字段
}

func main() {
	s := Student{
		Person: Person{"mike", 'm', 19},
		int:    18,
		mystr:  "hehehe",
	}
	fmt.Printf("s = %+v\n", s)
}
