package main

import (
	"fmt"
)

/*
	学生管理系统函数版
	具有能够查看，新增，删除学生
*/
type student struct {
	id   int64
	name string
}

var allStudent map[int64]*student //声明变量

func main() {
	//初始化变量
	allStudent = make(map[int64]*student)
	//打印菜单，等待用户选择，执行对应的函数
	for true {
		fmt.Println("欢迎光临学生信息管理系统！！！")
		fmt.Println(`
		1. 查看所有学生信息
		2. 新增学生
		3. 删除学生
		4. 退出系统
	`)
		fmt.Print("请输入你要选择的功能: ")
		var choice int
		fmt.Scanln(&choice)
		if choice == 1 || choice == 2 || choice == 3 {
			fmt.Printf("你选择了%d这个选项！\n", choice)

			switch choice {
			case 1:
				showAllStudent()
			case 2:
				addStudent()
			case 3:
				deleteStudent()
			}
		} else if choice == 4 {
			fmt.Println("系统已退出！")
			break
		} else {
			fmt.Println("你选择了无效的选项，请重新选择！")
		}
	}
}

// 把所有的学生信息都打印出来
func showAllStudent() {
	for k, v := range allStudent {
		fmt.Printf("学号： %d， 姓名：%s\n", k, v.name)
	}
}

//newStudent 是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

//向allStudent中添加一个学生
func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Println("请输入学生的学号:")
	fmt.Scanln(&id)
	fmt.Println("请输入学生的姓名:")
	fmt.Scanln(&name)
	//构造一个学生,将其添加到allStudent中去
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

//将一个学生从allStudent中删除
func deleteStudent() {
	var (
		deleteID int64
	)
	fmt.Println("请输入要删除的学生学号: ")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)
}
