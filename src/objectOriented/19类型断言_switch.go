package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3) //定义一个容量为3的空接口类型的切片
	i[0] = 1                    //int
	i[1] = "hello go"
	i[2] = Student{
		name: "tony",
		id:   888}
	//通过switch来实现类型断言
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int,内容为%d\n", index, value)
		case string:
			fmt.Printf("x[%d] 类型为string,内容为%s\n", index, value)
		case Student:
			fmt.Printf("x[%d] 类型为Student,内容为%+v\n", index, value)
		}
	}
}
