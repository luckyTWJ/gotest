package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"名字"`
	Age  int    `info:"age" doc:"年龄"`
}

func main() {

	var r resume
	findTag(&r)
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println("info:", field.Tag.Get("info"))
		fmt.Println("doc:", field.Tag.Get("doc"))
	}

}
