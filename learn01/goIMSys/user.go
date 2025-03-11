package main

import (
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建用户的API
func newUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	//启动监听当前User channel消息的gorutine
	go user.ListenMessage()
	return user
}

// 监听当前user channel 的方法 ,有消息就发送给客户端
func (this *User) ListenMessage() {

	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线功能
func (this *User) Online() {

	// 用户上线 加入到OnlineMap
	this.server.mapLock.Lock()
	this.server.OnlineUserMap[this.Name] = this
	this.server.mapLock.Unlock()
	//广播用户上线消息
	this.server.BroadCast(this, "on line!!!")
}

// 用户下线功能
func (this *User) Offline() {
	// 用户下线 OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineUserMap, this.Name)
	this.server.mapLock.Unlock()
	//广播用户上线消息
	this.server.BroadCast(this, "off line!!!")
}

// 给当前user对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理消息
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		//查询当前在线用户
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineUserMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "online...\n"
			//this.C <- onlineMsg
			//this.conn.Write([]byte(onlineMsg))
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//消息格式: rename|新用户名
		newName := strings.Split(msg, "|")[1]
		//newName := msg[7:]
		_, ok := this.server.OnlineUserMap[newName]
		if ok {
			this.SendMsg("The current user name is in use, please re-enter !\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineUserMap, this.Name)
			this.server.OnlineUserMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("You have updated your username ->" + newName + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//消息格式: to|张三|消息内容
		//获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("Message format is incorrect, please use\"to|zhang3|content\"format\n")
			return
		}
		//根据用户名得到对方user对象
		remoteUser, ok := this.server.OnlineUserMap[remoteName]
		if !ok {
			this.SendMsg("The user name does not exist! \n")
		}
		//	获取消息内容发送给对方 remoteUser
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("Message format is incorrect, please use\"to|zhang3|content\"format\n")
			return
		}
		if content == "offline" {
			this.Offline()
			remoteUser.SendMsg(this.Name + " said:" + content + "\n")
			return
		}
		remoteUser.SendMsg(this.Name + " said:" + content + "\n")

	}

	//广播消息
	//this.server.BroadCast(this, msg)
}
