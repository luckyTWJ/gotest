package zinface

import "net"

// 链接信息
type IConnection interface {
	//	启动链接 让当前链接开始工作
	Start()
	//	停止链接 结束当前链接工作
	Stop()
	//	获取当前链接的绑定socket conn
	GetTCPConnection() *net.TCPConn
	//	//	获取当前链接模块的链接id
	GetConnID() uint32
	//	获取远程客户端的TCP状态 IP port
	RemoteAddr() net.Addr
	//发送数据
	SendMsg(data []byte) error

	//该链接处理的方法router
	//GetRouter() IRouter
}

// 定义一个处理链接业务的方法
// []byte 接收数据的内容
// int 读取数据的长度
type HandleFunc func(*net.TCPConn, []byte, int) error
