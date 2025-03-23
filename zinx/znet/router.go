package znet

import "gotest/zinx/zinface"

/*
*
实现Router时 先嵌入 Router抽象类，然后根据需要对这个类进行重写
*/
type BaseRouter struct {
}

// baserouter方法都为空  是因为有的router不希望 有PreHandle PostHandle 业务
// 处理业务之前的钩子方法 Hook
func (router *BaseRouter) PreHandle(request zinface.IRequest) {

}

// 处理conn业务 的钩子方法
func (router *BaseRouter) Handle(request zinface.IRequest) {

}

// 处理conn业务之后的钩子方法
func (router *BaseRouter) PostHandle(request zinface.IRequest) {

}
