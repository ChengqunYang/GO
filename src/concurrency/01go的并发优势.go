package main

/*
	go从语言层面就支持了并发,同时并发程序的内存管理有时候是非常复杂的,
而go提供了自动垃圾回收机制
	go语言为并发编程而内置的上层API基于CSP(communicating sequential process,顺序通信模型)模型,
这就意味着显示锁都是可以避免的,因为go语言通过安全的通道发送和接受数据以实现同步
这大大简化了并发程序的编写
	一般情况下,一个普通的桌面计算机跑十几二十个线程就有点负载过大了,但是同样这台计算机可以
轻松的让成百上千个goroutine进行资源竞争
*/
func main() {
	/*
			goroutine是什么:
				goroutine是go并行设计的核心,goroutine说到底就是协程,它比线程更小,十几个goroutine可能体现在底层就是五六个线程,
		go语言内部帮你实现了这些goroutine之间的内存共享,执行goroutine只需要极少的栈内存(大概是4-5kb),当然会根据响应的数据伸缩,
		也正因为如此,可同时运行成千上万个并发任务,goroutine比thread更易用,更高效,更轻便.
	*/

	/*
			只需在函数调用的语句前加go关键字,就可以创建并发执行单元,开发人员无需了解任何执行细节,
		调度器会自动将其安排到合适的系统线程上执行
			在并发编程里,我们通常将一个过程切分成几块,然后让每个goroutine各自负责一块工作,当一个程序启动时,
		其主函数即在一个单独的goroutine中运行,我们叫它main goroutine(主协程),新的goroutine会用go语句来创建
	*/
}
