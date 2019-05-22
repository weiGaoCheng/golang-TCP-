package main

import (
	"fmt"
	"time"
	"net"
)

//模拟客户端
func main(){
	fmt.Println("client start...")

	time.Sleep(1*time.Second)
	//直接connect 服务器得到一个 已经建立好的conn句柄

	conn,err:=net.Dial("tcp","127.0.0.1:8999")
	if err != nil {
		fmt.Println("net.Dail error", err)
		return
	}
	for{
		_,err:=conn.Write([]byte("hello Zinx..."))
		if err != nil {
			fmt.Println("conn.Write error", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 *time.Second)
	}

	}
