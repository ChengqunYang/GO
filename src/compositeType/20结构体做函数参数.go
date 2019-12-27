package main

import "fmt"

//结构体做函数参数是值传递
type student struct {
	name string
	age  int
}

func update(s student)  {
	s.name = "王麻子"
}

func up(s *student) {
	s.name = "王麻子"
}
func main() {
	s:=student{name:"赵六",age:22}
	fmt.Println(s)
	update(s)   //值传递,形参无法修改实参
	fmt.Println(s)
	fmt.Println("--------------")
	up(&s)    //引用传递,形参可以修改实参
	fmt.Println(s)
}
