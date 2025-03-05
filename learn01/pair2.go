package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book2 struct {
}

func (this *Book2) ReadBook() {
	println("ReadBook")
}
func (this *Book2) WriteBook() {
	println("WriteBook")
}

func main() {
	b := Book2{}
	var r Reader
	r = &b
	r.ReadBook()
	fmt.Println("------------------")

	var w Writer
	//w = &b

	//r:pair<type:Book2,value:book{}地址>
	w, ok := r.(Writer) //断言 此处的w,r 具体的type是一致的
	fmt.Println("ok->", ok)
	fmt.Println("w->", w)

	fmt.Printf("w->%t\n", w)
	w.WriteBook()

}
