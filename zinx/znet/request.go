package znet

import (
	"gotest/zinx/zinface"
)

type Request struct {
	// 已经和客户端建立好链接
	conn zinface.IConnection
	//客户端请求的数据
	//data []byte

	msg zinface.IMessage
}

func (r *Request) GetConnection() zinface.IConnection {

	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {

	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
