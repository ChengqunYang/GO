package main

import "fmt"

func main() {
	//增加和修改
	m1:=map[int]string{1:"mike",2:"yoyo"}
	fmt.Println(m1)
	//修改键为2 对应的值
	m1[2] = "tony"
	fmt.Println(m1)
	//追加值,map底层自动扩容,和append类似
	m1[3] = "go"
	fmt.Println(m1)

	//map的遍历
	//第一个返回值为key,第二个返回值为value,遍历的结果是无序的
	for i, s := range m1 {
		fmt.Printf("%d====>%s\n",i,s)
	}
	//如何判断一个key是否存在
	//第一个返回值为key所对应的value,第二个返回值为key是否存在的条件,存在ok为true
	value,ok:= m1[4]
	if ok==true {
		fmt.Printf(value)
	}else {
		fmt.Println("key不存在")
	}
	//删除某一个key
	//删除key为1的内容
	delete(m1,1)
	fmt.Println(m1)
}
