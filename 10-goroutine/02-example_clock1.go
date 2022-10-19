package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 这是一个clock服务器，它会向连接的客户端输出服务器时间
// 这个服务器同一时间只能处理一个客户端，如果第2个客户端试图连接，必须等第1个断开后
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
