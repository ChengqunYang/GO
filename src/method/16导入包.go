package main

import (
	"fmt"
	"os"
)

/*import "fmt"
import "os"
import . "fmt"//调用函数,无需再通过包名
import io "fmt" //给fmt这个包起了个别名:io
import _ "fmt"   //忽略此包;有时,用户可能需要导入一个包,但是并不需要引用这个包的标识符,存在这种情况,可以使用空白标识符来重命名这个导入

*/

func main() {
	fmt.Println("aaa")
	list := os.Args
	fmt.Println(list)
}
