package main

import (
	"fmt"
)

//学生信息管理系统
/*
	有一个物件:
		1. 它保存了一些数据  ---> 结构体的字段
		2. 它有三个功能      --->结构体的方法
*/

type student struct {
	id   int64
	name string
}

//造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

var (
	allStudent map[int64]student
)

//查看学生
func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	//根据用户输入的内容,创建一个学生,将该学生加入到allStudent这个map中
	var (
		stuID   int64
		stuName string
	)
	fmt.Println("请输入用户的学号:")
	fmt.Scanln(&stuID)
	fmt.Println("请输入用户的姓名:")
	fmt.Scanln(&stuName)

	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功!")
}

//修改学生
func (s studentMgr) editStudent() {
	fmt.Println("请输入要修改的学生的学号: ")
	var stuID int64
	fmt.Scanln(&stuID)
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("你要修改的学生不存在!")
		return
	}
	fmt.Printf("你要修改的学生信息如下:  学号: %d 姓名:%s\n", stuObj.id, stuObj.name)
	fmt.Println("请输入学生的新名字:")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuID] = stuObj
	fmt.Println("修改成功!")
}

//删除学生
func (s studentMgr) deleteStudent() {
	fmt.Println("请输入要删除的学生的学号:")
	var stuID int64
	fmt.Scanln(&stuID)
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("你要删除的学生不存在!")
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("删除成功!")
}
