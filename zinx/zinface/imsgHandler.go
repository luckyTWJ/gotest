package zinface

//消息管理抽象层

type IMsgHandler interface {
	//调度对应的Router处理方法
	DoMsgHandler(request IRequest)
	//为消息添加具体的处理逻辑
	//根据msgID来绑定对应的处理逻辑
	AddRouter(msgId uint32, router IRouter)
	//StartOneWorker(workerId int, conn IConnection)
	//StartWorkerPool()
	//SendMsgToTaskQueue(request IRequest)
}
