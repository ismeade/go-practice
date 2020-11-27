package main

import (
	"fmt"
	"net"
)

func main() {
	// 1. 与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed, err", err)
		return
	}
	defer conn.Close()
	// 2. 发送数据
	conn.Write([]byte("HelloWorld"))
}
