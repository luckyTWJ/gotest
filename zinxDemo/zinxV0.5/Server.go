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
	fmt.Println("Call Router Handle...")
	// 先读取客户端的数据，再回写
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println(">>>>>>>call back ping err>>>>>>>>", err)

	}

}

func main() {
	//1.创建server句柄 使用zinx的api
	server := znet.NewServer("[zinx v0.5]")
	//2.给当前的server注册自定义的router
	server.AddRouter(&PingRouter{})

	//3.启动server
	server.Serve()
}
