package network

import (
	"fmt"
	"log"
	"net"
	"time"
)

// ConnManager 保存map映射
type ConnManager struct {
	Listen      net.Listener
	LocalIP     string
	ServerPeers map[string]*ServerPeer
}

// NewConnManager 新建一个ConnManager
func NewConnManager(ip string) *ConnManager {
	serverPeers := make(map[string]*ServerPeer)
	connManager := &ConnManager{
		LocalIP:     ip,
		ServerPeers: serverPeers,
	}
	return connManager
}

// Start 封装了startListen和acceptConn函数
func (connManager *ConnManager) Start(localIP string) {
	connManager.startListen(localIP)
	go connManager.acceptConn()
}

// startListen 开始监听某个端口
func (connManager *ConnManager) startListen(localIP string) bool {
	connManager.LocalIP = localIP
	fmt.Println(localIP)
	listen, err := net.Listen("tcp", connManager.LocalIP)
	if err != nil {
		log.Printf("开始监听时出现错误错误：%v", err)
		return false
	}
	connManager.Listen = listen
	fmt.Println("监听成功")
	return true
}

// RequestConn 开始发起一个连接，如果连接发送不成功就for循环一直发，直到发送成功break出循环
func (connManager *ConnManager) requestConn(targetIP string) *ServerPeer {
	serverPeer, ok := connManager.ServerPeers[targetIP]
	if ok {
		return serverPeer
	}
	var conn net.Conn
	var err error
	for {
		conn, err = net.DialTimeout("tcp", targetIP, time.Second)
		if err != nil {
			//log.Printf("建立连接不成功：%v", err)
		} else {
			break
		}
	}
	log.Printf("成功建立连接：%s", conn.RemoteAddr().String())
	ip := conn.RemoteAddr().String()
	newPeer := NewPeer(ip, conn)
	newServerPeer := NewServerPeer(newPeer)
	connManager.ServerPeers[ip] = newServerPeer //将新建连接加入到map中
	return newServerPeer
}

// acceptConn 接收一个连接,无限for循环
func (connManager *ConnManager) acceptConn() {
	for {
		log.Println("正在Accept")
		conn, err := connManager.Listen.Accept()
		if err != nil {
			log.Printf("接收连接时发生错误：%v", err)
			return
		}
		log.Printf("成功建立连接：%s", conn.RemoteAddr().String())
		ip := conn.RemoteAddr().String()
		newPeer := NewPeer(ip, conn)
		newServerPeer := NewServerPeer(newPeer)
		go newServerPeer.ReceiveMessage()
		connManager.ServerPeers[ip] = newServerPeer
	}
}

// BroadCast 向所有建立的连接广播消息
func (connManager *ConnManager) BroadCast(msg *Message) {
	for _, serverPeer := range connManager.ServerPeers {
		fmt.Println(serverPeer.peer.conn.LocalAddr().String())
		serverPeer.SendMessage(msg)
	}
}

//Send 封装了请求连接和连接成功之后的发送
func (connManager *ConnManager) Send(msg *Message, targetIP string) {
	serverPeer := connManager.requestConn(targetIP)
	serverPeer.SendMessage(msg)
}
