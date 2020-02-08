package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	//时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano()) //纳秒值

	//time.Unix()函数可以将时间戳转化为正常时间
	ret := time.Unix(1581134808, 0)
	fmt.Println(ret)
	fmt.Println(ret.Date())

	//时间间隔
	fmt.Println(time.Hour)
	fmt.Println(time.Minute)
	fmt.Println(time.Second)

	//时间操作  add,sub,equal,before,after
	fmt.Println(now.Add(24 * time.Hour))           //add
	fmt.Println(now.Sub(time.Unix(1581135306, 0))) //sub表示两个时间相减
	fmt.Println("_____________")
	fmt.Println(time.After(2 * time.Second))

	fmt.Println(now.Equal(time.Now())) //判断时间是否相同

	//定时器
	timer := time.Tick(3 * time.Second)
	var i int
	for t := range timer {
		fmt.Println(t) //1秒钟执行一次
		i++
		if i > 2 {
			break

		}
	}

	fmt.Println("__________")

	//格式化时间,把语言中的时间对象,转化为字符串类型的时间
	strtime := now.Format("2006-01-02  15:04:05")
	fmt.Println(strtime)
	fmt.Println("___________")
	//按照对应的格式解析字符串类型的时间返回
	timeObj, err := time.Parse("2006-01-02  15:04:05", strtime)
	if err != nil {
		fmt.Println("parse time failed,err:%v\n", err)
		return
	}
	fmt.Println(timeObj)

}
