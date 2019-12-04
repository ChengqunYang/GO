package main

import "fmt"

func main() {
	var a int = 10
	//每个变量有两层含义,变量的内存和变量的地址
	fmt.Printf("a = %d\n", a) //内存
	fmt.Printf("%v\n", &a)    //地址

	//保存某个变量的地址,需要使用指针类型, *int 保存 int的地址,**int 保存 *int的地址
	//声明(定义),定义只是特殊的声明
	//定义一个变量p,类型为*int
	var p *int
	p = &a
	fmt.Printf("p = %v,&a = %v\n", p, &a)

	//*p操作的不是P的内存,而是p指向的内存(就是a)
	*p = 666
	fmt.Printf("p = %v, a =%v\n", *p, a)
	fmt.Printf("p = %v,&a = %v\n", p, &a)

}
