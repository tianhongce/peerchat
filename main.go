package main

import (
	"fmt"

	"peerchat/network"
)

func main() {
	var server network.Server
	var localIP string
	fmt.Println("请输入本机IP地址：")
	fmt.Scanln(&localIP)
	serverPort := ":9999"
	server.StartServer(localIP + serverPort)
	server.Interaction()
}
