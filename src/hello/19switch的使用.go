package main

import "fmt"

func main() {
	var num int
	fmt.Println("请按下楼层:")
	fmt.Scan(&num)
	//switch支持一个初始化语句,初始化语句和变量本身,以分号分隔
	switch num = 2; num {
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
		fallthrough //不写默认为break,当前case做完后继续进行后面的case
	case 3:
		fmt.Println("按下的是3楼")
	case 4:
		fmt.Println("按下的是4楼")
		break
		//go语言保留了break关键字,跳出switch语句,不写默认包含break
	default:
		fmt.Println("发生错误了")
	}

	switch i := 5; i {
	case 1:
		fmt.Println("按下的是1")
	case 2:
		fmt.Println("按下的是2")
	case 3, 4, 5:
		fmt.Println("按下的yyy")
	}

	var score int
	fmt.Println("请输入分数:")
	fmt.Scan(&score)
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("一般")
	default:
		fmt.Println("及格")
	}
}
