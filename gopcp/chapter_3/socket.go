package socket

import (
	"fmt"
	"net"
	"strings"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8080"
	DLIMETER       = "\t"
)

func serverGo() {
	var l net.Listener
	l, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerLog("Listen error: %s", err)
		return
	}
	defer l.Close()
	printServerLog("Got listener. local address: %s", l.Addr())

	for {
		conn, err := l.Accept()
		if err != nil {
			printServerLog("Accept error: %s", err)
		}
		printServerLog("Got connection. Remote address: %s", conn.RemoteAddr())
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {}

func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args...)
}

func printClientLog(format string, sn int, args ...interface{}) {
	printLog("Client", sn, format, args...)
}
