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

// Test Handle
func (this *PingRouter) Handle(request zinface.IRequest) {
	fmt.Println("Call PingRouter Handle...")
	// 先读取客户端的数据，再回写
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println(">>>>>>>call back ping err>>>>>>>>", err)

	}

}

type HelloRouter struct {
	znet.BaseRouter
}

func (this *HelloRouter) Handle(request zinface.IRequest) {
	fmt.Println("Call HelloRouter Handle...")
	// 先读取客户端的数据，再回写
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("hello ...zinx...\n"))
	if err != nil {
		fmt.Println(">>>>>>>call back ping err>>>>>>>>", err)
	}

}

// 创建链接之后的钩子函数
func DoConnBegin(conn zinface.IConnection) {
	fmt.Println("======> DoConnBegin is Called ... ")
	err := conn.SendMsg(200, []byte("DoConnBegin..."))
	if err != nil {
		fmt.Println("SetOnConnStart SendMsg err:", err)
	}
}

// 链接断开之后的钩子函数
func DoConnLost(conn zinface.IConnection) {

	fmt.Println("DoConnLost is Called ... ")
}

func main() {
	//1.创建server句柄 使用zinx的api
	server := znet.NewServer("[ZINX SERVER]")

	//2.注册 链接的钩子函数
	server.SetOnConnStart(DoConnBegin)
	server.SetOnConnStop(DoConnLost)

	//2.给当前的server注册自定义的router
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloRouter{})

	//

	//3.启动server
	server.Serve()
}
