package main

import (
	"fmt"
	"gotest/zinx/zinface"
	"gotest/zinx/znet"
)

//使用zinx框架开发服务器端

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request zinface.IRequest) {
	fmt.Println("Call Router PreHandle...")

	//err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping..."))

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...ping...ping...\n"))
	if err != nil {
		fmt.Println(">>>>>>>call back before ping err>>>>>>>>", err)

	}

}

// Test Handle
func (this *PingRouter) Handle(request zinface.IRequest) {
	fmt.Println("Call Router Handle...")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Handle ping...ping...ping...\n"))
	if err != nil {
		fmt.Println(">>>>>>>call back Handle ping err>>>>>>>>", err)

	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request zinface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...ping...ping...\n"))
	if err != nil {
		fmt.Println(">>>>>>>call back PostHandle ping err>>>>>>>>", err)

	}

}

func main() {
	//1.创建server句柄 使用zinx的api
	server := znet.NewServer("[zinx v0.4]")
	//2.给当前的server注册自定义的router
	server.AddRouter(&PingRouter{})

	//3.启动server
	server.Serve()
}
