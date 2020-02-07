package main

import "fmt"

//在通常情况下,向程序使用方报告错误状态的方式可以是返回一个额外的error类型
/*
	但是,当遇到不可恢复的错误状态的时候,如数组访问越界,空指针引用等,这些运行时错误会引起panic异常,
这时,上述错误的处理方式显然就不适合了,反过来讲,在一般情况下,我们不应通过panic函数来报告普通的错误,
而应该只把它作为报告致命错误的一种方式,当某些不应该发生的场景发生时,我们就应该调用panic
一般而言,当panic异常发生时,程序会中断运行,并立即执行在该goroutine,随后,程序崩溃并输出日志信息,
日志信息包括panicvalue和函数调用的的堆栈信息

不是所有的panic异常都来自运行时,直接调用内置的panic函数也会引发panic异常,panic函数接受任何值作为参数

*/

func testa() {
	fmt.Println("aaaaaaaaaaaaaaa")
}

func testb() {
	fmt.Println("bbbbbbbbbb")
	//显式调用panic函数,导致程序中断
	panic("this is a panic test")
}

func testc() {
	fmt.Println("ccccccc")
}

func testd() {
	fmt.Println("ddddd")
}
func main() {
	testa()
	testb() ////显式调用panic函数
	testc()
	testd()
}
