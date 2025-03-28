package znet

type Message struct {
	Id      uint32 // 消息的ID
	DataLen uint32 // 消息数据包长度
	Data    []byte // 消息内容
}

// 提供一个创建msg的方法
func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
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
