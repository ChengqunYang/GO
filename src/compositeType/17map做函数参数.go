package main

import (
	"fmt"
)

func del(m map[int]string)  {
	delete(m,1)
}
//map做函数参数是引用传递
func main() {
	m := map[int]string{1: "mike", 2: "yoyo", 3: "go"}
	fmt.Println(m)
	//在del函数内删除某一个key
	del(m)
	fmt.Println(m)
}
