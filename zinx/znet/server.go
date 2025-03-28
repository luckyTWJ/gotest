package znet

import (
	"errors"
	"fmt"
	"gotest/zinx/utils"
	"gotest/zinx/zinface"
	"net"
	"time"
)

// IServer接口的视线 服务器模块
type Server struct {
	Name      string
	IP        string
	Port      int
	IPVersion string
	// 路由 当前server添加一个router 处理
	//Router zinface.IRouter

	//当前server的消息管理模块，用来绑定msgId和对应的处理业务API关系
	MsgHandler zinface.IMsgHandler
	//server 的链接管理器
	ConnMgr zinface.IConnManager

	//该server创建链接之后自动调用Hook函数
	OnConnStart func(conn zinface.IConnection)

	//该server销毁之前自动调用的Hook函数
	OnConnStop func(conn zinface.IConnection)
}

func (s *Server) Start() {

	fmt.Printf("[Zinx] Server Name: %s, IP: %s ,Port %d, is starting....\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Println("[Zinx] Version:", utils.GlobalObject.Version,
		"MaxConn:", utils.GlobalObject.MaxConn,
		"WorkerPoolSize", utils.GlobalObject.WorkerPoolSize,
		"MaxWorkerTaskLen", utils.GlobalObject.MaxWorkerTaskLen)
	fmt.Println("===================================================")

	go func() {
		// 0.开启消息队列及worker工作池
		s.MsgHandler.StartWorkerPool()

		//	1.获取一个TCP Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.ResolveTCPAddr err:", err)
			return
		}

		//2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("net.ListenTCP err:", err)
			return
		}
		//defer listener.Close()
		fmt.Println("strat Zinx server ", s.Name, " success ,now listening...")
		//	3.阻塞等待的客户端链接 处理客户端链接业务
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err:", err)
				continue
			}
			fmt.Println("链接成功--cid->", cid)

			//设置最大连接数 超过最大值则关闭客户端连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Write([]byte("Too Many connections..."))
				fmt.Println("Too Many connections", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//defer conn.Close()
			//与客户端建立链接 绑定到自定义connection
			dealConnc := NewConnection(s, conn, cid, s.MsgHandler)
			cid++
			//启动当前链接的业务处理
			go dealConnc.Start()

		}
	}()

}

// 定义当前客户端所绑定的hanler api 目前写死
func CallbackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//fmt.Println("[Conn Handle] CallbackToClient...")

	fmt.Printf("come from client[%s] : %s, cnt = %d\n", conn.RemoteAddr(), data[:cnt], cnt)

	//	回显
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("CallbackToClient err")
	}
	return nil
}

func (s *Server) GetConnMgr() zinface.IConnManager {
	return s.ConnMgr
}

func (s *Server) Stop() {
	//服务器的资源 状态 停止或者回收
	fmt.Println("[Zinx] Server ", s.Name, " is stoping...")
	s.ConnMgr.ClearConn()
}
func (s *Server) Serve() {
	s.Start() //异步的
	//TODO 启动服务之后额外的业务

	//	阻塞 状态 否则注go退出 listener 的go也会退出

	for {
		time.Sleep(10 * time.Second)
	}

}
func (s *Server) AddRouter(msgId uint32, router zinface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("add router success...")
}

// 初始化server模块方法
func NewServer(name string) zinface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		//Router:    nil,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

// 注册OnConnectionStart钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(conn zinface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 注册OnConnectionStop钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(conn zinface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用OnConnectionStart钩子函数的方法
func (s *Server) CallOnConnStart(conn zinface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("--->Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

// 调用OnConnectionStop钩子函数的方法
func (s *Server) CallOnConnStop(conn zinface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("--->Call OnConnStop()...")
		s.OnConnStop(conn)
	}

}
