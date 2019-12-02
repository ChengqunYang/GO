package main

func test01() int {
	//函数被调用时,x才分配空间,才初始化为0
	var x int //没有初始化,值为零
	x++
	return x * x //函数调用完毕,x就会自动释放
}

//函数的返回值是一个匿名函数,返回的是个函数类型(无参,一个int型的返回值 )
func test02() func() int {
	var x int //没有初始化,值为0
	return func() int {
		x++
		return x * x //函数调用完毕,x就会自动释放
	}
}
func main() {
	/*  值都为0
	println(test01())
	println(test01())
	println(test01())
	println(test01())*/

	//返回值为一个匿名函数,所以通过一个函数类型来接收,再通过f来调用闭包函数
	//闭包不关心这些捕获了的变量和常量是否已经超出了作用域
	//所以只要闭包还在使用它,这些变量就会存在
	f := test02()
	println(f())
	println(f())
	println(f())
	println(f())
}
