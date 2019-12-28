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
		name: "mike",
		id:   666}
	//通过类型断言来识别不同的类型
	for index, data := range i {
		//第一个返回的是切片元素,第二个返回判断结果的真假
		if _, ok := data.(int); ok == true {
			fmt.Printf("x[%d] 类型为int,内容为%d\n", index, data)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] 类型为string,内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d] 类型为Student,内容为%+v\n", index, value)
		}
	}
}
