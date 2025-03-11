package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP   string
	Port int

	//	message 广播管道
	Message chan string
	//在线用户列表
	OnlineUserMap map[string]*User
	mapLock       sync.RWMutex
}

func NewServer(IP string, Port int) *Server {
	server := &Server{
		IP:            IP,
		Port:          Port,
		OnlineUserMap: make(map[string]*User),
		//广播管道
		Message: make(chan string),
	}
	return server

}
func (this *Server) Handler(conn net.Conn) {
	//	当前链接的业务
	//fmt.Println("conn success !!! info:", conn)

	user := newUser(conn, this)
	user.Online()
	//监听用户是否活跃的channel
	isLive := make(chan bool)
	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for true {
			n, err := conn.Read(buf)
			if n == 0 {
				//this.BroadCast(user, "off line!!")
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			// 接收客户端发送的消息 去除\r\n
			msg := string(buf[:n-1])
			//	消息处理
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	//	当前handler阻塞
	for true {
		select {
		case <-isLive:
			//	当前用户是活跃的，应该重置定时器
			//不做任何事情，为了激活select,更新select
		case <-time.After(time.Second * 900):
			//	已经超时
			user.SendMsg("Time out! please log in again!\n")
			//	销毁资源
			close(user.C)
			//	关闭连接
			conn.Close()
			//	退出当前handler
			return
		}

	}

}

// 监听Message广播消息channel的 gooroutine 一旦有消息就发送给全部在线的User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		//	将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, cliUser := range this.OnlineUserMap {
			cliUser.C <- msg
		}
		this.mapLock.Unlock()
	}

}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}
func (this *Server) Start() {
	//1 socket listen...
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//4.close listen socket...
	defer listener.Close() //close listen socket

	//启动监听Message的goroutine
	go this.ListenMessager()

	for {
		//2.accept...
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		//3.do handler
		go this.Handler(conn)
	}

	//fmt.Printf("Server is starting on %s:%d\n", this.IP, this.Port)
}
