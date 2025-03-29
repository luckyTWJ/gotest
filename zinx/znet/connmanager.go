package znet

import (
	"fmt"
	"gotest/zinx/zinface"
	"sync"
)

/*
链接管理模块
*/

type ConnManager struct {
	// 管理链接的集合
	connections map[uint32]zinface.IConnection
	//保护链接集合 读写锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]zinface.IConnection),
	}

}

// 添加连接
func (cm *ConnManager) Add(conn zinface.IConnection) {
	//保护链接集合 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	// 将conn加入到链接集合中
	cm.connections[conn.GetConnID()] = conn
	fmt.Println("connID = ", conn.GetConnID(), " add to ConnManager Success... 链接数量：", cm.Len())
}

// 删除连接
func (cm *ConnManager) Remove(conn zinface.IConnection) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	delete(cm.connections, conn.GetConnID())
	fmt.Println("connID = ", conn.GetConnID(), " remove to ConnManager Success... 链接数量：", cm.Len())

}

// 根据connID获取链接
func (cm *ConnManager) Get(connID uint32) (zinface.IConnection, error) {
	//	加读锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("connID = %d is not exist", connID)
	}

}
func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

// 清除 所有链接
func (cm *ConnManager) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for connID, conn := range cm.connections {
		//停止
		conn.Stop()
		//删除
		delete(cm.connections, connID)
	}
	fmt.Println("Clear All Connections Success!!!", cm.Len())

}
