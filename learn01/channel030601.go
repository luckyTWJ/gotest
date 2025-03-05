package main

import (
	"fmt"
	"time"
)

func main() {
	//testchannel()

	//testchancache()
	//testClose()
	testrange()
	fmt.Println("main finished...")
}

func testrange() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭chan 只有确定不发送数据了 尝试关闭chan ；
		//如果关闭了 chan，那么chan就永远是关闭的，不能再发送数据了
		//关闭chan后可以继续 从chan中取值；
		close(c)
	}()

	//阻塞等待 从chan中取值 迭代从range读
	for data := range c {
		fmt.Println("data->", data)
	}

}
func testClose() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭chan 只有确定不发送数据了 尝试关闭chan ；
		//如果关闭了 chan，那么chan就永远是关闭的，不能再发送数据了
		//关闭chan后可以继续 从chan中取值；
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println("data->", data)
		} else {
			fmt.Println("channel is closed")
			break
		}
	}
}

// 缓存测试
func testchancache() {
	c := make(chan int, 3) //带有缓存的chan
	fmt.Println("len(c)->", len(c), "cap(c)->", cap(c))

	go func() {
		defer fmt.Println("子go结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go正在运行，发送的元素=", i, "len(c)->", len(c), "cap(c)->", cap(c))
		}
	}()
	//当子go运行完成之后，再从chan中取值
	time.Sleep(2 * time.Second)
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("main goroutine 正在运行，接收的元素=", num, "len(c)->", len(c), "cap(c)->", cap(c))
	}

	//当channel满时，再往chan中发送数据，会阻塞
	//当channel空时，再从chan中取数据，会阻塞

}
func testchannel() {
	//	定义一个chan
	c := make(chan int)
	go func() {
		defer fmt.Println("goroutine end...")
		fmt.Println("gorooutine running...")
		c <- 666 //发送数据给c
	}()
	//阻塞等待 num值
	num := <-c
	fmt.Println("main goroutine end...", num)
}
