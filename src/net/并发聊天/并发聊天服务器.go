package main

//主要功能,转发广播
import (
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()
	//创建一个结构体,用户名和地址默认一样
	cli := Client{make(chan string), cliAddr, cliAddr}
	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个协程,专门给当前客户端发送信息
	go WriteMsgToClint(cli, conn)

	//广播某个人在线
	//message <- "{" + cli.Addr + "}" + cli.Name + ":  longin"
	message <- MakeMsg(cli, "login")
	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here")

	hasData := make(chan bool) //定义一个对方是否有数据发送的通道
	isQuit := make(chan bool)  //定义一个接受用户退出的通道

	//新建一个协程,接受用户发送过来的数据
	go func() {
		buf := make([]byte, 2048)
		for true {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开,或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-1])
			//查看当前在线的用户都有哪些
			if len(msg) == 3 && msg == "who" {
				//遍历map,给当前用户发送所有在线成员
				conn.Write([]byte("user list:"))
				for _, user := range onlineMap {
					msg = user.Addr + ":" + user.Name
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//给用户重命名
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok"))

			} else {
				//转发此内容
				message <- MakeMsg(cli, msg)
			}
			hasData <- true //能进入,即表示有数据写入
		}
	}()

	for true {
		//通过select检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)           //将当前用户从map中移除
			message <- MakeMsg(cli, "login out") //广播某个用户下线
			return
		case <-hasData:
		case <-time.After(60 * time.Second): //60秒后给管道中写入数据
			delete(onlineMap, cliAddr)                    //将当前用户从map中移除
			message <- MakeMsg(cli, "time out leave out") //广播某个用户超时并断开连接
			return
		}
	}
}

func WriteMsgToClint(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)
	for true {
		msg := <-message //没有消息前,阻塞
		//遍历map,给map中的每个成员都发送此消息
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "{" + cli.Addr + "}" + cli.Name + ":" + msg
	return
}

type Client struct {
	C    chan string //用户发送数据的管道
	Name string      //用户名
	Addr string      //网络地址
}

//保存在线用户
var onlineMap map[string]Client

//创建一个通信的管道
var message = make(chan string)

func main() {
	//监听
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	//新开一个协程,用来转发消息,只要有消息来,就遍历map,给每个成员都发送这个消息
	go Manager()

	//主协程,循环阻塞,等待用户连接
	for true {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}
}
