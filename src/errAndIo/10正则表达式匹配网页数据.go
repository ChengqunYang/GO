package main

import (
	"fmt"
	"regexp"
)

func main() {
	//反引号 表示原生字符串
	buf := `
<html>
	<div class="col left-column" id="main-left-cloumn">
	<div class="tab"  id="cate0"><i class="fa fa-reorder"></i> 全部教程</div>
		<div class="sidebar-box gallery-list">
		 			<div class="design" id="cate1">
				<div class="navto-nav"><i class="fa fa-external-link"></i> HTML / CSS</div>
			</div>
					<div class="design" id="cate2">
				<div class="navto-nav"><i class="fa fa-external-link"></i> JavaScript</div>
			</div>
					<div class="design" id="cate3">
				<div class="navto-nav"><i class="fa fa-external-link"></i> 服务端</div>
			</div>
					<div class="design" id="cate4">
				<div class="navto-nav"><i class="fa fa-external-link"></i> 数据库</div>
			</div>
					<div class="design" id="cate5">
				<div class="navto-nav"><i class="fa fa-external-link"></i> 移动端</div>
			</div>
					<div class="design" id="cate6">
				<div class="navto-nav"><i class="fa fa-external-link"></i> XML 教程</div>
			</div>
		</div>
	</div>
	<div class="col middle-column-home">
	      <div class="codelist codelist-desktop cate1" >
         
      <h2><i class="fa fa-list"></i> HTML / CSS</h2>
   	
      	
       <a class="item-top item-1" href="//www.runoob.com/html/html-tutorial.html"><h4>【学习 HTML】</h4>
	</html>
`
	reg := regexp.MustCompile(`<div.*>(?s:(.*))</div>`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}

	//提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	//fmt.Println("result = ", result)

	//过滤<div></div>
	for _, text := range result {
		//	fmt.Println("text[0] =" ,text[0])
		fmt.Println("text[1] = ", text[1])
	}
}
