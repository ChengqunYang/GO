package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//Person类型,实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name = %s,sex = %c,age = %d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生,继承person字段,成员和方法都继承了
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

//Student也实现了一个方法,这个方法和Person方法同名
func (tmp *Student) PrintInfo() {
	fmt.Println("Student tmp = ", *tmp)
}

func main() {
	s := Student{
		Person: Person{"mike", 'm', 20},
		id:     12,
		addr:   "bj"}
	s.PrintInfo()        //调用的是student的方法,就近原则(先找本作用域的方法,找不到再找继承的方法)
	s.Person.PrintInfo() //显式调用继承来的方法
}
