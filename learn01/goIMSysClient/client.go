package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

var ServerIp string
var ServerPort int

// 初始化命令行参数
func init() {
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "设置服务器ip")
	flag.IntVar(&ServerPort, "port", 8888, "设置服务器端口")
}
func NewClient(serverIp string, serverPort int) *Client {

	//创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		Name:       "Client",
		flag:       9999,
	}
	//链接服务器
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return nil
	}

	client.conn = conn
	//返回对象
	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>>>请输入合法范围内的数字>>>>>>>>")
		return false
	}

}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}

		switch client.flag {
		case 1:
			fmt.Println(">>>>>>>公聊模式>>>>>>>>")
			client.PublicChat()
			break
		//
		case 2:
			fmt.Println(">>>>>>>私聊模式>>>>>>>>")
			client.PrivateChat()
			break
		case 3:
			fmt.Println(">>>>>>>更新用户名>>>>>>>>")
			client.updateName()
			break

		}
	}
}

// 处理server回应的消息 直接显示到终端
func (client *Client) dealResponse() {
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞等待
	io.Copy(os.Stdout, client.conn)
}
func (client *Client) PublicChat() {
	//提示用户输入消息
	fmt.Println(">>>>>>>请输入聊天内容，exit退出.")

	//	发送给服务器
	var chatMsg string
	//fmt.Println(&chatMsg)
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) > 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println(">>>>>>>sendMsg err>>>>>>>>", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>>>>>>请输入聊天内容，exit退出.")

		fmt.Scanln(&chatMsg)
	}
}

func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(">>>>>>>who err>>>>>>>>", err)
		return
	}

}
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	client.SelectUsers()

	fmt.Println(">>>>>>>请输入聊天对象[用户名]>>>>>>>> exit退出。")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println(">>>>>>>请输入聊天内容>>>>>>>> exit退出。")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) > 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println(">>>>>>>sendMsg err>>>>>>>>", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>>>>>>请输入聊天内容>>>>>>>> exit退出。")
			fmt.Scanln(&chatMsg)

		}
		client.SelectUsers()
		fmt.Println(">>>>>>>请输入聊天对象[用户名]>>>>>>>> exit退出。")
		fmt.Scanln(&remoteName)

	}
}
func (client *Client) updateName() bool {
	fmt.Println(">>>>>>>请输入用户名>>>>>>>>")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(">>>>>>>rename err>>>>>>>>", err)
		return false
	}
	return true

}
func main() {
	// 解析命令行参数
	flag.Parse()
	client := NewClient(ServerIp, ServerPort)
	if client == nil {
		fmt.Println(">>>>>>>link server err!!>>>>>>>>")
		return
	} else {
		fmt.Println(">>>>>>>link server success!!>>>>>>>>")
	}
	go client.dealResponse()
	// 启动客户端业务
	client.Run()

	//go func() {
	//	for {
	//		// 在这里添加你的业务逻辑
	//	}
	//}()
	// 防止程序立即退出
	//for {
	//	time.Sleep(time.Second)
	//}
	// 等待中断信号
	//select {}
}
