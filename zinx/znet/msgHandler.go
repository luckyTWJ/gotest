package znet

import (
	"fmt"
	"gotest/zinx/utils"
	"gotest/zinx/zinface"
)

type MsgHandler struct {
	// 存放每个MsgID对应的处理方法
	Apis map[uint32]zinface.IRouter
	// 负责Worker取任务的消息队列
	//TaskQueue []chan zinface.IRequest
	// 业务工作Worker池的数量
	//WorkerPoolSize uint32

	//	worker工作池数量
	WorkerPoolSize uint32
	// 消息队列 数量 = workerPoolSize
	TaskQueue []chan zinface.IRequest
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]zinface.IRouter),
		TaskQueue:      make([]chan zinface.IRequest, utils.GlobalObject.WorkerPoolSize), //从全局配置中获取
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
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

// 启动一个worker工作池 (开启工作池的动作只能发生一次，只能有一个worker工作池)
func (mh *MsgHandler) StartWorkerPool() {
	//根据workerPoolSize 分别开启worker  每个worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//1 创建一个worker，然后把worker通过taskQueue传给worker
		//2 worker从taskQueue中取出一个request，执行router中的业务
		mh.TaskQueue[i] = make(chan zinface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		//	启动当前的worker 阻塞等待消息从channel传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个worker工作流程
func (mh *MsgHandler) StartOneWorker(workerID int, taskQueue chan zinface.IRequest) {
	fmt.Println("StartOneWorker()-->workerID = ", workerID, " is started...")
	//	不断阻塞等待对应的队列消息
	for {
		select {
		//如果有消息过来，则从taskQueue中取出消息
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}

}

// 将消息交给taskQueue，进行出列，并执行
func (mh *MsgHandler) SendMsgToTaskQueue(request zinface.IRequest) {
	//	平均分配给不同的worker
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnID(), " request msgID=", request.GetMsgID(), " to workerID=", workerID)
	//将消息发送给worker对应的taskQueue
	mh.TaskQueue[workerID] <- request
}
