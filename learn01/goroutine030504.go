package main

import (
	"fmt"
	"time"
)

func newTask() {

	i := 0
	for {
		i++
		fmt.Printf("new Goroutine :i=%d\n", i)
		time.Sleep(1 * time.Second)

	}
}
func main() {
	//创建一个goroutine 执行newTask()流程
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine :i=%d\n", i)
		time.Sleep(1 * time.Second)

	}
	//主goroutine退出 其他从 全部退出
	fmt.Println("main goroutine exit")
}
