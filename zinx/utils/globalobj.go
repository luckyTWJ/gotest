package utils

import (
	"encoding/json"
	"fmt"
	"gotest/zinx/zinface"
	"io/ioutil"
	"os"
)

/*
*存储一切有关zinx全局参数 供其他模块使用
一些参数是通过zinx.json由用户进行配置
*/
type GlobalObj struct {
	// Server
	TcpServer zinface.IServer //当前zinx全局的server对象
	Host      string          // 当前服务器主机IP
	TcpPort   int             // 当前服务器主机监听的tcp端口
	Name      string          // 当前服务器主机的名称

	// Zinx
	Version          string //当前zinx版本号
	MaxPacketSize    uint32 //都包最大尺寸
	MaxConn          int    //当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 //业务工作Worker池的数量
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 //SendBuffMsg发送消息的缓冲最大长度
}

// 从配置文件中加载一些用户自定义的参数
func (g *GlobalObj) Reload() {
	// 打印当前工作目录
	wd, _ := os.Getwd()
	fmt.Println("Current working directory:", wd)

	fmt.Println("Start loading zinx.json...")

	//data, err := ioutil.ReadFile("./zinx/zinxDemo/zinxV0.4/conf/zinx.json")
	data, err := ioutil.ReadFile("zinxDemo/zinxV0.4/conf/zinx.json")
	//将json文件解析到 GlobalObject
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 定义一个全局的对外GlobalObj
var GlobalObject *GlobalObj

// 提供一个init方法 初始化当前GlobalObj对象
// 只要导入util包 init方法会自动调用  在main函数之前调用
func init() {
	// 如果配置文件没有加载，默认使用init初始化一个对象
	GlobalObject = &GlobalObj{
		Host:             "0.0.0.0",
		TcpPort:          8999,
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		MaxConn:          1000,
		MaxPacketSize:    512,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    256,
	}
	// 从配置文件中加载一些用户自定义的参数
	//调试的时候注释掉
	GlobalObject.Reload()
}
