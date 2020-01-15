package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func DoWork(start int, end int) {
	fmt.Printf("正在爬取%d 到%d的页面", start, end)
	//明确目标(要知道你准备在哪个范围或者网站去搜索)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		fmt.Println("url = ", url)

		//爬取数据(将所有的网站的内容全部爬下来)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err = ", err)
			continue
		}
		//把内容写入文件
		fileName := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.create err1 = ", err1)
			continue
		}
		f.WriteString(result)
		f.Close() //关闭文件
	}
}

//爬取网页内容
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	//读取网页内容
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if n == 0 { //读取结束,或者,出问题
			fmt.Println("resp.body.read err = ", err)
			break
		}
		result += string(buf[:n])
	}
	return
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页(>=1): ")
	fmt.Scan(&start)

	fmt.Printf("请输入终止页(>=起始页): ")
	fmt.Scan(&end)

	DoWork(start, end)
}
