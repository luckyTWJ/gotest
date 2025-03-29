package zinface

//定义一个服务器接口

type IServer interface {
	Start()

	Stop()

	Serve()

	//	路由功能：给当前服务注册一个路由方法 ，供客户端的链接处理使用
	//AddRouter(router IRouter)
	AddRouter(msgId uint32, router IRouter)
	//获取当前server的链接管理器
	GetConnMgr() IConnManager

	//注册OnConnectionStart钩子函数的方法
	SetOnConnStart(func(conn IConnection))
	//注册OnConnectionStop钩子函数的方法
	SetOnConnStop(func(conn IConnection))

	//调用OnConnectionStart钩子函数的方法
	CallOnConnStart(conn IConnection)
	//调用OnConnectionStop钩子函数的方法
	CallOnConnStop(conn IConnection)
}
