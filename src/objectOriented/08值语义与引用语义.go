package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值

//接收者为普通变量,非指针,值语义其实是一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("&p = %p\n", &p)
}

//接收者为指针变量,引用传递,引用语义就是修改对应的地址的数据
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("&p = %p\n", p)
}
func main() {
	var s1 Person
	fmt.Printf("&s1 = %p\n", &s1)
	s1.SetInfoPointer("lisi", 'w', 20)
	s1.SetInfoValue("zhangsan", 'm', 23)
	fmt.Println(s1)
}
