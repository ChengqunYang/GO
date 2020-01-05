package main

/*
	通道默认是双向的,也就是既可以往里面发送数据,也可以从里面接受数据
	但是,我们常见的一个通道作为参数进行传递,而且希望对方是单向使用的
	要么只让他发送数据,要么只让他接受数据,这时,我们可以指定通道的方向

	<-chan   表示数据从管道出来,对应调用者就是得到通道的数据,当然就是输入
	chan<-	 表示数据进入管道,要把数据写进管道,对应调用者就是输入
*/
func main() {
	//创建一个channel,双向的
	ch := make(chan int)

	//双向的channel转化为单向的channel
	var writeCh chan<- int = ch //只能写,不能读

	var readCh <-chan int = ch //只能读,不能写

	//可以将channel隐式转换为单向队列,只收或者只发,不能将单向channel转化为管道
	writeCh <- 666
	<-readCh

	/*	ch2 := make(chan int) //因为类型不匹配无法将一个单向channel转化为一个普通的channel
		ch2 = writeCh*/

}
