package ziface


/*
  Server模块的抽象层
*/
type IServer interface {

	//开启服务器方法
	Start()
	//关闭服务器方法
	Stop()
	//开启业务服务方法
	Serve()

}
