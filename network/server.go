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

// StartServer 开启server，监听本地IP的连接
func (server *Server) StartServer(localIP string) {
	server.localIP = localIP
	server.connManager = NewConnManager(localIP)
	server.connManager.Start(localIP)
}

// Input 进行业务逻辑的交互
// 如果是广播消息，消息的格式：bd@text
// 如果是发送消息，消息的格式：send@x.x.x.x(IP地址)@text
func (server *Server) Input(text string) {

	textString := strings.Split(text, "@")
	switch textString[0] {
	case "bd":
		msg := Message{textString[1], time.Now().Format("2006-01-02 15:04:05"), server.localIP, server.targetIP}
		server.connManager.BroadCast(&msg)
	case "send":
		textIP := strings.Split(textString[1], "@")
		msg := Message{textIP[1], time.Now().Format("2006-01-02 15:04:05"), server.localIP, server.targetIP}
		server.connManager.Send(&msg, textIP[0])
	default:
		fmt.Println("Invalid command， input again")
	}
}

// Interaction 进行交互
// 使用for循环保证主线程一直在接收键盘的操作
func (server *Server) Interaction() {
	for {
		server.Input(InputMsg())
	}
}

// InputMsg 读取键盘上的输入
func InputMsg() string {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	if string(data) == "exit" {
		os.Exit(3)
	}
	return string(data)
}
