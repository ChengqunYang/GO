package main

func testa(x int) {
	var a [10]int
	a[x] = 111 //当x的值大于9时就会导致程序崩溃
}

func main() {
	//当发生数组越界异常时,程序自动调用panic函数
	testa(10)
}
