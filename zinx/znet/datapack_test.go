package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//测试文件以 test结尾  单元测试

func TestDataPack(f *testing.T) {
	/*
		模拟服务器发送数据
	*/
	//1.先创建一个socketTCP
	listener, err := net.Listen("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("start server failed, err:", err)
		return
	}
	//创建一个go 负责从客户端处理业务
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept failed, err:", err)
				return
			}

			go func(conn net.Conn) {
				//处理客户端请求
				//	---------------拆包过程------------
				// 1.创建一个封包对象
				dp := NewDataPack()
				for {
					//1.第一次从conn读 把包的head读出来
					headData := make([]byte, dp.GetHeadLen())
					if _, err := io.ReadFull(conn, headData); err != nil {
						fmt.Println("read head failed, err:", err)
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("unpack err, err:", err)
						return
					}
					if msgHead.GetDataLen() > 0 {
						// msg是有数据的 需要再次读取
						//2.第二次从conn读 根据head中的datalen 在读取datalen内容
						msg := msgHead.(*Message)                 // 类型断言
						msg.Data = make([]byte, msg.GetDataLen()) // 创建一个切片 大小为datalen

						if _, err := io.ReadFull(conn, msg.Data); err != nil {
							fmt.Println("read msg data failed, err:", err)
							return
						}
						fmt.Println("---->Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
					}

				}

			}(conn)
		}
	}()

	// 2.创建一个客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start failed, err:", err)
		return
	}
	// 3.创建一个封包对象
	dp := NewDataPack()
	// 模拟粘包过程 封装两个msg一起发送
	// 1.先发送一个msg1包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	// 2.再发一个msg2包
	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'g', 'o', 'o', 'd', 'b', 'y', 'e'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 failed, err:", err)
		return
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 failed, err:", err)
		return
	}
	// 将两个包粘在一起
	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)
	//defer conn.Close()
	//
	//客户端阻塞
	select {}
}
