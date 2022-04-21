package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	shouldReturn := checkError(err)
	if shouldReturn {
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("你叫什么名字呢？")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClientName := strings.Trim(clientName, "\r\n")
	fmt.Printf("原来你叫 %v 啊！\n", trimmedClientName)

	for {
		fmt.Printf("你想发送什么到服务器呢，%v? 不想发送的话就输入 q 退出吧！\n", trimmedClientName)
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClientName + " 说 " + trimmedInput))
	}
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
