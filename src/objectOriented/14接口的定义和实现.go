package main

import "fmt"

type Humaner interface {
	//方法只有声明没有实现,由别的类型实现
	sayhi()
}
type Student struct {
	name string
	id   int
}

//Studnet实现了此方法
func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s,%d] sayhi\n", tmp.name, tmp.id)

}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s,%s] sayhi\n", tmp.addr, tmp.group)

}

type Mystr string

//Mystr也实现了此方法
func (tmp *Mystr) sayhi() {
	fmt.Printf("Mystr[%s] sayhi\n", *tmp)
}

//定义一个普通函数,函数的参数为接口类型
//只有一个函数,却可以有不同的表现(多态)
func WhoSayHi(i Humaner) {
	i.sayhi()
}
func main() {
	//定义一个接口类型的变量
	var i Humaner
	//只要实现了此接口的类型,那么这个类型的变量就可以赋值给接口
	s := &Student{
		name: "mike",
		id:   666,
	}
	s.sayhi()
	i = s
	i.sayhi()

	t := Teacher{
		addr:  "bj",
		group: "go",
	}
	t.sayhi()

	i = &t
	i.sayhi()

	var str Mystr = "hello mike"
	i = &str
	str.sayhi()
	i.sayhi()

	fmt.Println("-------------------")
	//多态,不同的对象,对于同一个消息的响应不同
	WhoSayHi(s)
	WhoSayHi(&t)
	WhoSayHi(&str)

	fmt.Println("--------------")
	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = &t
	x[2] = &str
	//第一个返回下标,第二个返回下标对应的值
	for _, i := range x {
		i.sayhi()
	}
}
