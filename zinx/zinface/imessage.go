package zinface

/*
将请求的消息  封装在 Message中
*
*/
type IMessage interface {
	// 获取消息的ID
	GetMsgID() uint32
	// 获取消息的长度
	GetDataLen() uint32
	// 获取消息的内容
	GetData() []byte
	// 设置消息的ID
	SetMsgID(uint32)
	// 设置消息的内容
	SetData([]byte)
	// 设置消息的长度
	SetDataLen(uint32)
}
