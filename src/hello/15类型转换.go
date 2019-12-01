package main

import "fmt"

func main() {
	//不兼容类型
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)
	//bool类型不能转化为整型
	//fmt.Printf("flag = %d\n",int(flag))
	//0就是假,非0就是真
	//整型也不能转化为bool类型
	//flag = bool(1)

	//字符型本质上就是整型
	var ch byte
	ch = 'a'
	fmt.Printf("ch = %c\n", ch)

	var i int
	i = int(ch)
	fmt.Printf("%d", i)

}
