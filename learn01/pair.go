package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//testPair()
	fmt.Println("-------------------")

	testOpenFile()

}

func testPair() {
	var pair string
	//类型 静态的类型
	//pair<statictype:string,value:'hello'>
	pair = "hello"

	//类型
	var allType interface{}
	//pair<statictype:string,value:'hello'> 类型传递
	allType = pair
	//断言 是否string类型
	str, ok := allType.(string)
	fmt.Println(str)
	fmt.Println("ok->", ok)
	if ok {
		fmt.Println("ok true->", ok)
	} else {
		fmt.Println("ok false->", ok)
	}
}
func testOpenFile() {
	// 使用 os.OpenFile 而不是 os.Open
	//tty:pair<statictype:*os.File,value:'/dev/tty 文件描述符'>
	var tty, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error")
		fmt.Println(err)
	} else {
		fmt.Println(tty)
		fmt.Println("open file success")
	}
	defer tty.Close()
	var r io.Reader
	//r:pair<statictype:*os.File,value:'/dev/tty 文件描述符'>
	r = tty
	var w io.Writer
	w = r.(io.Writer)
	w.Write([]byte("hello this is test!"))

}
