package zinface

// 将客户端请求的链接信息和数据封装在 一个Requerst中
type IRequest interface {
	//获取请求连接信息
	GetConnection() IConnection
	//获取请求消息的数据
	GetData() []byte
	//获取请求的消息ID
	GetMsgID() uint32
}
