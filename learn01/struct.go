package main

import "fmt"

// 声明数据类型myint
type myint int
type Book struct {
	title string
	auth  string
}

func main() {
	var a myint = 10
	fmt.Println("a=", a)
	fmt.Printf("a type is %T\n", a)

	var b Book
	b.title = "go语言"
	b.auth = "golang"
	fmt.Println("b=", b)
	fmt.Printf("b type is %T\n", b)
	chageBook(b)
	fmt.Println(" change b=", b)
	changeBook(&b)
	fmt.Println(" change b=", b)

}

func chageBook(book Book) {
	//传递book的副本
	book.auth = "ttt"
}
func changeBook(book *Book) {
	//传递book的指针
	book.auth = "ttt*"
}
