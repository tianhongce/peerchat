package main

import (
	"bufio"
	"fmt"
	"os"

	"chat_v3.2/network"
)

func main() {
	var server network.Server
	var targetPort string
	var localPort string
	fmt.Println("请输入监听的端口：")
	fmt.Scanln(&localPort)
	fmt.Println("请输入连接的端口：")
	fmt.Scanln(&targetPort)
	localIP := "127.0.0.1:"
	server.StartServer(localIP + localPort)
	go server.Interaction(InputMsg(), localIP+targetPort)
}

// InputMsg 信息输入
func InputMsg() string {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	if string(data) == "exit" {
		os.Exit(3)
	}
	return string(data)
}
