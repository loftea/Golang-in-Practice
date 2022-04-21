package main // 学习 net.Dial() 方法的使用

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.0.32.10:80")
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "192.0.32.10:80")
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80")
	checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("在连接时发生错误：%v", err)
		os.Exit(1)
	}
	fmt.Printf("连接成功，通过 %v\n", conn)
}
