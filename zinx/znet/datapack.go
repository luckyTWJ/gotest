package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gotest/zinx/utils"
	"gotest/zinx/zinface"
)

/*
*
封包拆包模块
直接面向TCP连接中的数据流，把TCP的数据流封装成包，在拆包时，
按照 zinx 的格式拆包，再把包在封装到 Message 中，然后传递给 api层
*/
type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取消息头长度
func (dp *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4字节) + ID uint32(4字节)
	return 8
}

// 封包方法，创建一个msg的data数据，并把msg的id，dataLen，data all into the []byte
func (dp *DataPack) Pack(msg zinface.IMessage) ([]byte, error) {
	//创建一个存放msg的byteBuffer
	dataBuff := bytes.NewBuffer([]byte{})
	// 写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	// 写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	// 写data
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

// 拆包方法 分两次读，第一次读dataLen,第二次从datalen偏移量开始读  将包的二进制数据转成msg
func (dp *DataPack) Unpack(binaryData []byte) (zinface.IMessage, error) {
	//
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	msg := &Message{}
	// 读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	// 判断dataLen的长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {

		return nil, errors.New("too large msg data received!")
	}
	return msg, nil
}
