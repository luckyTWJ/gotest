package zinface

type IDataPack interface {
	// 获取消息头长度
	GetHeadLen() uint32
	// 封包方法，创建一个msg的data数据，并把msg的id，dataLen，data all into the []byte
	Pack(msg IMessage) ([]byte, error)
	// 拆包方法，将包的二进制数据转成msg
	Unpack(binaryData []byte) (IMessage, error)
}
