package main

import (
	"fmt"

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
	server.Interaction(localIP + targetPort)

}
