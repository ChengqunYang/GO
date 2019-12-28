package main

import "fmt"

//打印函数的参数就是一个可变参数的空接口类型(可以接受任意个数,任意类型的参数)
type kong interface {
}

func main() {
	//空接口,万能类型,他可以保存任意类型的值
	var i interface{} = 1
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)

	//定义在外面的空接口
	var k kong = 2
	fmt.Println(k)
	k = "abcd"
	fmt.Println(k)

}
