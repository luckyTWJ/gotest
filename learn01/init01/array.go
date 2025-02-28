package main

import (
	"fmt"
)

func main() {
	//固定长度数组  值copy
	var arr [10]int
	arr2 := [10]int{1, 2, 3, 4}
	arr4 := [4]int{11, 22, 33, 44}
	for i := 0; i < len(arr); i++ {
		arr[i] = i
		fmt.Println(arr[i])
	}

	//动态数组
	arr5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("arr5 type=%T\n", arr5)
	/*for j := range arr2 {
		fmt.Print(arr2[j])
	}*/
	for index, value := range arr2 {
		fmt.Println("index->", index, "  value->", value)
	}

	//查看数组的数据类型
	fmt.Printf("arr type=%T\n", arr)
	fmt.Printf("arr2 type=%T\n", arr2)

	printArr(arr4)
	fmt.Println("========")
	printArrs(arr5)
	fmt.Println("========")
	printArrs(arr5)
}

// 固定数组  值传递
func printArr(a [4]int) {
	for index, val := range a {
		fmt.Println("index->", index, ",value->", val)
	}
	//固定数组   值copy 原始数字值不变
	a[0] = 111
}

// 动态数组 引用传递
func printArrs(a []int) {
	//_,匿名 不使用
	for _, val := range a {
		fmt.Println("value->", val)
	}

	//动态数组   引用地址 值会变
	a[0] = 111100
}
