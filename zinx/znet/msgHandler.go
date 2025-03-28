package znet

import (
	"fmt"
	"gotest/zinx/zinface"
)

type MsgHandler struct {
	// 存放每个MsgID对应的处理方法
	Apis map[uint32]zinface.IRouter
	// 负责Worker取任务的消息队列
	//TaskQueue []chan zinface.IRequest
	// 业务工作Worker池的数量
	//WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]zinface.IRouter),
	}
}
func (mh *MsgHandler) DoMsgHandler(request zinface.IRequest) {
	// 判断是否由该msg对应的router
	router, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is not found!")
		return
	}
	// 根据msgID来获取对应的router
	//router := mh.Apis[request.GetMsgID()]
	// 执行对应router的Handle方法
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

// 为消息添加具体的处理逻辑
// 根据msgID来绑定对应的处理逻辑
func (mh *MsgHandler) AddRouter(msgId uint32, router zinface.IRouter) {
	if mh.Apis == nil {
		mh.Apis = make(map[uint32]zinface.IRouter)
	}
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeat register msgId" + string(msgId))
	}
	mh.Apis[msgId] = router
	fmt.Println("add api msgID = ", msgId, " success!!")
}
