package zinface

/*
*
路由接口抽象
路由里的数据都是IRequest
*/
type IRouter interface {
	//处理业务之前的钩子方法 Hook
	PreHandle(request IRequest)
	//处理conn业务 的钩子方法
	Handle(request IRequest)
	//处理conn业务之后的钩子方法
	PostHandle(request IRequest)
}
