package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}
func main() {
	//顺序初始化,每个成员必须初始化
	var s1 Student = Student{1,"张三",'m',18,"bj"}
	fmt.Println(s1)

	//指定初始化内容,没有初始化的成员,自动赋值为0
	s2:=Student{name:"王五",addr:"xa"}
	fmt.Println(s2)

	//指针类型的初始化
	s3:=&Student{name:"李四",addr:"wh"}
	fmt.Printf("tyep s3 is %T\n",s3)
	fmt.Println(*s3)
}
