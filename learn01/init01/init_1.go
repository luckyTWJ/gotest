package main

import (
	"fmt"
	_ "time"
	//_ "/init01/lib2/lib2"
	"gotest/learn01/init01/lib1"
)

func main() {
	fmt.Println("8888")
	// 首字母大写 对外开放 否则包内调用
	lib1.Lib1Test()
	//myLib2.Lib2Test()
	//Lib1Test()
	var a int = 10
	var b int = 20
	//swap(a, b)
	swapAddr(&a, &b)
	fmt.Println("a=", a, " b=", b)

	var p *int //一级指针
	p = &a
	fmt.Println("p->", p)
	fmt.Println("a->", &a)
	var pp **int = &p // 二级指针
	fmt.Println("pp->", &pp)

}

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

// 地址交换
func swapAddr(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
