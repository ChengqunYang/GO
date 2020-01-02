package main

/*
	用于在多个协程之间共享内存
	goroutine运行在相同的地址空间,因此访问共享内存必须做好同步,goroutine奉行
通过通信来共享内存,而不是共享内存来通信
	引用类型channel是csp模式的具体实现,用于多个goroutine通讯,其内部实现了同步,确保并发安全
*/
/*
		和map类似,channel也对应着一个make创建的底层数据结构的引用
	当我们复制一个channel或者用于函数参数传递时,我们只是拷贝了一个channel引用,因此调用者和被调用者将引用同一个channel对象,
	和其他的类型相同,channel的零值是nil
*/
func main() {
	/*
		channel的创建:
		make(chan Type,0)  type为channel的值的类型
		make(chan Type,capacity)
		当capacity等于0时,channel是无缓冲阻塞读写的,当capacity>0时,channel有缓冲是非阻塞的
		直到写满capacity个元素才阻塞写入
	*/

}
