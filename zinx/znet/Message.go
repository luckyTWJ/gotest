package znet

type Message struct {
	Id      uint32 // 消息的ID
	DataLen uint32 // 消息数据包长度
	Data    []byte // 消息内容
}

func (msg *Message) GetMsgID() uint32 {
	return msg.Id
}
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}
func (msg *Message) GetData() []byte {
	return msg.Data
}
func (msg *Message) SetMsgID(id uint32) {
	msg.Id = id
}
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}
