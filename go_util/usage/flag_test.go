package usage

import (
	"flag"
	"fmt"
	"testing"
)

/*
go语言 flag库是用于解析命令行参数的标准库，使用该库可以方便地对程序的参数进行解析，从而实现灵活的程序设计。

使用场景：
	1.程序需要接收多个命令行参数，并据此改变其行为；
	2.程序需要在不同的运行模式之间切换；
	3.程序需要在运行时自动检测并加载配置文件等。

使用方法：
	1.安装flag库，使用 import "flag" 导入；
	2.在代码中定义需要接受的命令行参数，使用 flag.TypeVar(&value, "name", defaultValue, "description") 定义；
	3.在代码中使用 flag.Parse() 对命令行参数进行解析；
	4.在命令行中传入参数时，使用 -name=value 的格式进行传参。
*/

func TestFlag(t *testing.T) {
	var name string
	// var age int
	// var isMale bool
	flag.StringVar(&name, "name", "Tom", "name of person")
	flag.Parse()
	fmt.Printf("name: %v\n", name)
}
