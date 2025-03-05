package main

import "fmt"

// interface{} 空接口 任意类型 万能数据类型
type Book1 struct {
	auth string
}

func main() {
	book := Book1{"GOlang"}
	myFunc(book)

	myFunc(111199900)

}

// interface{} 空接口 任意类型 万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called....")
	fmt.Println(arg)
	//如何区分底层数据类型
	//	给interface{} 提供‘断言’机制
	switch arg.(type) {
	case string:
		fmt.Println("arg is string type")
	case int:
		fmt.Println("arg is int type")
	default:
		fmt.Println("unknow type")
	}

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type", value)
	}

	fmt.Println("========")

}
