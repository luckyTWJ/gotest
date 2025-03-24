package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	//	1.链接远端服务器 得到conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	//	链接调用write写数据
	for {
		_, err := conn.Write([]byte("hello zinx v0.4...."))
		if err != nil {
			fmt.Println("write conn err, exit!", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read buf err, exit!", err)
			return
		}
		fmt.Printf("server call back : %s, cnt = %d\n", buf, cnt)
		//cpu阻塞
		time.Sleep(3 * time.Second)
	}

}
