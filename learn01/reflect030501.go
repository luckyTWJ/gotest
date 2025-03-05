package main

import (
	"fmt"
	"reflect"
)

// valueOf 获取变量的value
// typeOf 获取变量的type
func reflect1() {
	var a int = 10
	fmt.Println(a)

	v := reflect.ValueOf(a)
	fmt.Println(v)

	t := reflect.TypeOf(a)
	fmt.Println(t)

	fmt.Println(v.Int())

	fmt.Println(t.Name())

	fmt.Println(t.Kind())

	fmt.Println(v.Kind())

	fmt.Println(v.CanSet())
}

func reflectNum(arg interface{}) {
	fmt.Println("type->", reflect.TypeOf(arg))
	fmt.Println("value->", reflect.ValueOf(arg))
}
func main() {
	//reflectNum("100")
	//reflect1()

	DoFileAndMethod(User{Name: "zhangsan", Age: 18, ID: "100"})
}

type User struct {
	Name string
	Age  int
	ID   string
}

func DoFileAndMethod(input interface{}) {
	//获取input 的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType)
	fmt.Println("inputType.Kind() is :", inputType.Kind())
	//	获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	//通过type获取里面的字段
	//1.获取interface的reflect.Type 通过Type获取里面的NumField字段 进行遍历
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("field name is %s,field type is %s,field offset is %d\n", field.Name, field.Type, field.Offset)
		fmt.Printf(" %s:  %v =%v\n", field.Name, field.Type, value)
	}
	fmt.Println("----------------")
	for i := 0; i < inputValue.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("method name is %s,method type is %s,method offset is %d\n", method.Name, method.Type, method.Index)
		fmt.Printf("%s:%v\n", method.Name, method.Type)
	}
}
