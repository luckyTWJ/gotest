package main

import (
	"fmt"
	"gotest/zinx/znet"
	"io"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client0 start...")
	time.Sleep(1 * time.Second)
	//	1.链接远端服务器 得到conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	count := 1
	//	链接调用write写数据
	for {
		count++
		//发送封包的msg消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMessage(10, []byte("Zinx V0.7 client0 Test "+string(count)+" Message...")))
		if err != nil {
			fmt.Println("client pack err: ", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("client write err: ", err)
			return
		}

		//服务器应该回复我们一个message数据 msgId：1
		//先读取流中的head部分 得到ID和dataLen
		//在根据dataLen 第二次读取data数据
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("client read head err: ", err)
			return
		}
		//将二进制的head拆包到 msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack err: ", err)
			return
		}
		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message) //类型断言 转换
			msg.Data = make([]byte, msg.GetDataLen())
			//读到msg.Data中
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("client read data err: ", err)
				return
			}
		}
		fmt.Println("client receive msg: id->", msgHead.GetMsgID(), " Len->", msgHead.GetDataLen(), "data->", string(msgHead.GetData()))
		//cpu阻塞
		time.Sleep(10 * time.Second)
	}

}
