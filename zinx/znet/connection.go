package znet

import (
	"errors"
	"fmt"
	"gotest/zinx/utils"
	"gotest/zinx/zinface"
	"io"
	"net"
)

type Connection struct {
	// 当前连接的socket tcp套接字
	Conn *net.TCPConn
	// 当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool
	// 该连接的处理方法  API
	//handleAPI zinface.HandleFunc
	//告知当前链接已退出、停止 chan
	exitChan chan bool // 由reader告知当前链接已经退出的channel

	//该链接处理的方法router
	//Router zinface.IRouter

	MsgHandler zinface.IMsgHandler

	//	用于无缓冲的管道 用于读 写两个goroutine之间的通信
	msgChan chan []byte
}

//初始化链接模块的方法

func NewConnection(conn *net.TCPConn, connID uint32, msgHandler zinface.IMsgHandler) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI: callback_api,
		MsgHandler: msgHandler,
		msgChan:    make(chan []byte),
		exitChan:   make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running...]")

	defer fmt.Println("ConnID = ", c.ConnID, "[ Reader is exit ] ,remote addr is ", c.RemoteAddr().String())
	//如果出现异常，关闭链接
	defer c.Stop()

	for {

		//创建一个拆包解包对象
		dp := NewDataPack()
		//读取客户端msg Head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read head err:", err)
			c.exitChan <- true //通知writer
			break
		}
		fmt.Println("--headData-->", headData)
		//拆包，得到msgID和msgData 放在msg消息中
		//msgHead, err := dp.Unpack(headData)
		//将二进制的head拆包到 msg结构体中
		msg, err := dp.Unpack(headData)

		if err != nil {
			fmt.Println("unpack err:", err)
			break
		}
		fmt.Println("unpack success: ", msg)
		//根据dataLen 再次读取Data 放在msg.Data中
		if msg.GetDataLen() > 0 {
			//msg := msgHead.(*Message) //类型断言 转换

			data := make([]byte, msg.GetDataLen())

			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error:", err)
				break
			}
			msg.SetData(data)

		}
		fmt.Println("服务器收到 拆包后的msgID:", msg.GetMsgID(), "dataLen:", msg.GetDataLen(), "data:", string(msg.GetData()))

		//得到当前conn数据的Rquerst请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}
		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经启动工作池机制，将消息交给Worker处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从路由中 找到注册绑定的Conn对应的router调用
			go c.MsgHandler.DoMsgHandler(&req)
		}

		//go func(req zinface.IRequest) {
		//	c.MsgHandler.DoMsgHandler(req)
		//
		//}(&req)

		//	调用当前链接的API处理方法  callback 回调
		//if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		//	fmt.Println("handleAPI err:", err)
		//	break
		//}
	}
}

/*
写消息goroutine 专门发送给客户端消息的模块
*/
func (conn *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running...]")
	defer fmt.Println("ConnID = ", conn.ConnID, "[ Writer is exit ] ,remote addr is ", conn.RemoteAddr().String())

	//不断的阻塞等待channel消息 发送给客户端
	for {
		select {
		case data, ok := <-conn.msgChan:
			if ok {
				//有数据要写给客户端
				if _, err := conn.Conn.Write(data); err != nil {
					fmt.Println("Send Data error:, ", err, " Conn Writer exit")
					return
				}
			}
		case <-conn.exitChan:
			//此时reader已退出 关闭Writer
			return
		}
	}

}

// 启动链接 让当前链接开始工作
func (conn *Connection) Start() {
	fmt.Println("ConnID = [", conn.ConnID, "] is Start working...")

	// 读写分离  开两个goroutine  一个读取客户端发送数据  一个负责写数据
	go conn.StartReader()

	//启动写消息的goroutine
	go conn.StartWriter()

}

// 停止链接 结束当前链接工作
func (conn *Connection) Stop() {
	fmt.Println("ConnID = ", conn.ConnID, " stop! ")
	if conn.isClosed == true {
		return
	}
	conn.isClosed = true
	conn.Conn.Close()
	//通知writer退出
	conn.exitChan <- true
	//回收资源
	close(conn.exitChan)
	close(conn.msgChan)
}

// 获取当前链接的绑定socket conn
func (conn *Connection) GetTCPConnection() *net.TCPConn {
	return conn.Conn
}

// //	获取当前链接模块的链接id
func (conn *Connection) GetConnID() uint32 {
	return conn.ConnID
}

// 获取远程客户端的TCP状态 IP port
func (conn *Connection) RemoteAddr() net.Addr {
	return conn.Conn.RemoteAddr()
}

// 发送数据

// 提供一个setMsg方法 将我们要发给客户端的数据 先进行封包 在发送
func (conn *Connection) SendMsg(msgId uint32, data []byte) error {

	if conn.isClosed == true {
		return errors.New("connection closed when set msg")
	}

	//将msg进行封包 msgdataLen + msgID + data
	dp := NewDataPack()
	//msgDataLen|msgID|data 格式
	pack, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id = ", msgId)
		return errors.New("pack error msg ")
	}
	/*if _, err := conn.Conn.Write(pack); err != nil {
		fmt.Println("write msg error msg id = ", msgId)
		return errors.New("conn Write error")
	}*/
	// 将pack的数据发送给客户端
	conn.msgChan <- pack
	return err

}
