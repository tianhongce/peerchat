package network

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Server 封装connManager
type Server struct {
	connManager *ConnManager
	localIP     string
	targetIP    string
}

// StartServer 开启服务，监听……，开启连接……
func (server *Server) StartServer(localIP string, targetIP string) {
	server.localIP = localIP
	server.targetIP = targetIP
	server.connManager = NewConnManager(localIP)
	server.connManager.Start(localIP, targetIP)
}

// Input 进行业务逻辑的交互
func (server *Server) Input(text string) {

	ts := strings.Split(text, "@")
	msg := Message{ts[1], time.Now().Format("2006-01-02 15:04:05"), server.localIP, server.targetIP}
	switch ts[0] {
	case "bd":
		server.connManager.BroadCast(&msg)
	case "send":
		server.connManager.Send(&msg, server.targetIP)
	default:
		fmt.Println("Invalid command")
	}

}

// Interaction 进行交互
func (server *Server) Interaction() {
	for {
		server.Input(InputMsg())
	}
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
