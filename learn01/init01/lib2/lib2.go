package lib2

import "fmt"

func init() {
	fmt.Println("lib2 init()...")
}

// 首字母大写 对外开放 否则包内调用
func Lib2Test() {
	fmt.Println("lib2Test() ...")
}
