package main

import (
	"fmt"
	"os"
)

//学生信息管理系统
var smr studentMgr //声明一个全局的变量学生管理对象:smr

//菜单函数
func showMenu() {
	fmt.Println(`
	1. 查看所有学生
	2. 添加学生
	3. 修改学生
	4. 删除学生
	5. 退出`)
}
func main() {
	fmt.Print("Welcome ssm!")
	smr = studentMgr{allStudent: make(map[int64]student, 100)}
	for true {
		showMenu()
		//等待用户输入
		fmt.Print("请输入要选择的功能的序号:")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 || choice == 2 || choice == 3 || choice == 4 || choice == 5 {
			fmt.Println("你输入的是:", choice)
			switch choice {
			case 1:
				smr.showStudents()
			case 2:
				smr.addStudent()
			case 3:
				smr.editStudent()
			case 4:
				smr.deleteStudent()
			case 5:
				fmt.Println("系统已退出!!!")
				os.Exit(0) //0表示正常退出
			}
		} else {
			fmt.Println("你的输入有误,请重新输入:")
		}

	}
}
