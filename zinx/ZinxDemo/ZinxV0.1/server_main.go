package main

import "zinx/znet"

func main() {
	//创建一个server对象
	s := znet.NewServer("ZinxV0.1")
	//让server对象启动服务器
	s.Serve()
}
