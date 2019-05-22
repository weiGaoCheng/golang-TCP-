package znet

import (
	"zinx/ziface"
	"net"
	"fmt"
	"io"
)

/*
Server模块的实现层
*/
type Server struct {
	//服务器的名称
	Name string
	//tcp4  or other
	IPVersion string
	//服务器绑定的ip地址
	IP string
	//服务器绑定的端口
	Port int
}

//初始化的new方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}

//开启服务器方法
func (s *Server) Start() {
	fmt.Printf("【start] Server Linstenner at IP :%s, Port :%d, is starting..\n", s.IP, s.Port)

	//1 创建套接字  ：得到一个TCP的addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error", err)
		return
	}

	//2 监听服务器地址
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("Listen  TCP error", err)
		return
	}

	//3 阻塞等待客户端发送请求，
	//起一个go程去循环
	go func() {
		for {
			//阻塞等待客户端请求,
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Listen  TCP error", err)
				continue
			}
			//此时conn就已经和对端客户端连接
			go func() {
				//4 客户端有数据请求，处理客户端业务(读、写)
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil || err == io.EOF {
						fmt.Println("conn read error", err)
						break
					}
					fmt.Printf("recv client buf %s, cnt = %d\n", buf, n)
					//写
					_,err=conn.Write(buf[:n])
					if err!=nil{
						fmt.Println("conn write error",err)
						continue
					}
				}

			}()

		}
	}()

}

//关闭服务器方法
func (s *Server) Stop() {

}

//开启业务服务方法
func (s *Server) Serve() {
	//启动server的监听功能
	s.Start()

	//防止服务器读写完后退出
	select {}
}
