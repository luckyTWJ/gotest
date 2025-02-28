package lib1

import "fmt"

func init() {
	fmt.Println("lib1 init()...")
}

// 首字母大写 对外开放 否则包内调用
func Lib1Test() {
	fmt.Println("lib1Test() ...")
}
