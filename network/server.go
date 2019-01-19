package network

import (
	"strings"
	"time"
)

// Server 封装connManager
type Server struct {
	connManager *ConnManager
}

// StartServer 开启服务，监听……，开启连接……
func (server *Server) StartServer(localIP string) {
	server.connManager.Start(localIP)

}

// Interaction 进行业务逻辑的交互
func (server *Server) Interaction(text string, targetIP string) {
	ts := strings.Split(text, "@")
	msg := Message{ts[1], time.Now().Format("2006-01-02 15:04:05"), server.connManager.LocalIP, targetIP}
	switch ts[0] {
	case "bd":
		server.connManager.ServerPeer.Broadcast(&msg)
	case "send":
		server.connManager.RequestConn(targetIP)
		server.connManager.ServerPeer.SendMessage(&msg)
	}
}
