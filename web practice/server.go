package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("请输入 Host 和 port！")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Acceept:")
		go connHandler(conn)
	}

}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")

	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")

	println("监听：", listener.Addr().String())

	return listener
}

func connHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("连接至： ", connFrom)
	sayHello(conn)

	for {
		var ibuf []byte = make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(length, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT
		}

	}
DISCONNECT:
	err := conn.Close()
	println("关闭连接：", connFrom)
	checkError(err, "Close: ")

}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "写错误："+string(wrote)+"字节。")
}

func handleMsg(length int, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		println(">")
	}
}

func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}
