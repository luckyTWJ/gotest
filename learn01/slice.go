package main

import "fmt"

func main() {
	//申明数组 并且初始化
	slice1 := []int{1, 2, 3}

	fmt.Println("hello world")

	fmt.Printf("len slice1= %d,slice = %v\n", len(slice1), slice1)
	//申明切片 没有给slice分配空间
	var slice2 []int
	slice2 = make([]int, 3)
	slice2[1] = 199
	fmt.Printf("len slice2= %d,slice = %v\n", len(slice2), slice2)
	//声明切片 同时分配空间
	//	容量是5 长度是3
	var slice3 []int = make([]int, 3, 5)
	fmt.Printf("len slice3= %d,cap = %d,slice = %v\n", len(slice3), cap(slice3), slice3)
	for i := 0; i < len(slice3); i++ {
		slice3[i] = i
		fmt.Println(slice3[i])
	}
	slice3 = append(slice3, 666, 777, 999)

	fmt.Printf("len slice3= %d,cap = %d,slice = %v\n", len(slice3), cap(slice3), slice3)
	//左闭右开 [1,3)
	slice4 := slice3[1:3]

	fmt.Printf("len slice4= %d,cap = %d,slice = %v\n", len(slice4), cap(slice4), slice4)

	//	切片属于 引用类型
	slice5 := make([]int, 5)
	//将slice3中的值 copy到slice5中
	copy(slice5, slice3)
	fmt.Printf("len slice5= %d,cap = %d,slice = %v\n", len(slice5), cap(slice5), slice5)

}
