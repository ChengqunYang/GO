package main

import "fmt"

//有参无返回值函数的定义,普通参数列表(固定参数)
func MyFunc01(a int) {
	a = 111
	fmt.Println("a = ", a)
}
func MyFunc02(a int, b int) {
	fmt.Printf("a = %d,b = %d\n", a, b)
}
func MyFunc03(a, b int, c string, d, e float64) {
	//a,b为int,c为string,d,e为float64
}

//有参无返回值函数的定义,不定参数列表
func MyFunc04(args ...int) { //...int ...type不定参数类型
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

func MyFunc05(temp ...int) {
	//将temp这个不定参数全部传递给MyFunc05这个函数
	MyFunc04(temp...)
	/*
		//将temp这个不定参数的从2开始的参数传递给MyFunc05这个函数
		MyFunc04(temp[2:]...)
		//将temp这个不定参数的从0开始到2(不包含)的参数传递给MyFunc05这个函数
		MyFunc04(temp[:2]...)
	*/
}

func main() {
	//调用有参无返回值的函数
	MyFunc01(666)
	MyFunc02(3, 5)
	MyFunc04(2, 3, 54)
	MyFunc05(2, 34, 5, 7, 2)
}
