package znet

import (
	"errors"
	"fmt"
	"gotest/zinx/utils"
	"gotest/zinx/zinface"
	"net"
)

type Connection struct {
	// 当前连接的socket tcp套接字
	Conn *net.TCPConn
	// 当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool
	// 该连接的处理方法  API
	handleAPI zinface.HandleFunc
	//告知当前链接已退出、停止 chan
	exitChan chan bool

	//该链接处理的方法router
	Router zinface.IRouter
}

//初始化链接模块的方法

func NewConnection(conn *net.TCPConn, connID uint32, r zinface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI: callback_api,
		Router:   r,
		exitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("StartReader() is running...")

	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer fmt.Println("ConnID = ", c.ConnID, " Reader is exit ,remote addr is ", c.RemoteAddr().String())
	//如果出现异常，关闭链接
	defer c.Stop()

	for {
		//	读取客户端发送的数据到buf 最大512字节
		buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			c.exitChan <- true
			continue
		}
		//得到当前conn数据的Rquerst请求数据
		req := Request{
			conn: c,
			data: buf,
		}

		//、从路由中 找到注册绑定的Conn对应的router调用

		go func(req zinface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)

		//	调用当前链接的API处理方法  callback 回调
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("handleAPI err:", err)
		//	break
		//}
	}
}

// 启动链接 让当前链接开始工作
func (conn *Connection) Start() {
	fmt.Println("ConnID = [", conn.ConnID, "] is Start working...")

	// 读写分离  开两个goroutine  一个读取客户端发送数据  一个负责写数据
	go conn.StartReader()

	for {
		select {
		case <-conn.exitChan:
			return
		}

	}
}

// 停止链接 结束当前链接工作
func (conn *Connection) Stop() {
	fmt.Println("ConnID = ", conn.ConnID, " stop! ")
	if conn.isClosed == true {
		return
	}
	conn.isClosed = true
	conn.Conn.Close()
	//conn.exitChan <- true
	//回收资源
	close(conn.exitChan)
}

// 获取当前链接的绑定socket conn
func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}

// //	获取当前链接模块的链接id
func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}

// 获取远程客户端的TCP状态 IP port
func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

// 发送数据
func (conn *Connection) SendMsg(data []byte) error {
	if conn.isClosed == true {
		return errors.New("connection closed when send msg")
	}
	_, err := conn.Conn.Write(data)
	if err != nil {
		fmt.Println("SendMsg err:", err)
		return errors.New("connection closed when send msg")
	}
	return nil
}
