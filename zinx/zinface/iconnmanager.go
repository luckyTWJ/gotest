package zinface

/*
链接管理模块
*/
type IConnManager interface {
	//添加连接
	Add(conn IConnection)
	// 删除连接
	Remove(conn IConnection)
	// 根据connID获取链接
	Get(connID uint32) (IConnection, error)
	Len() int
	ClearConn()
}
