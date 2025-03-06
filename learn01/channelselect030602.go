package main

import "fmt"

// 单流程下的go只能监控一个channel状态，select可以完成多个channel状态的监控
func main() {

	c1 := make(chan int)
	quit := make(chan int)

	//	 sub go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println("sub goroutine 正在运行，接收的元素=", <-c1, "len(c1)->", len(c1), "cap(c1)->", cap(c1))
		}
		//发送数据到quit中
		quit <- 0

	}()

	//main go
	fibnc(c1, quit)

}

func fibnc(c, quit chan int) {
	x, y := 1, 1
	for {
		//select 监控 chan 状态
		select {
		case c <- x:
			//如果c有空位，则执行case1    6666
			x, y = y, x+y

		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
