package main

import "gotest/zinx/znet"

//使用zinx框架开发服务器端

func main() {
	//1.创建server句柄 使用api
	server := znet.NewServer("[zinx v0.2]")
	//2.启动server
	server.Serve()
}
