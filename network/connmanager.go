package network

import (
	"log"
	"net"
	"time"
)

// ConnManager 负责监听的开始，发送接收消息时被调用
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

// Start 封装了startListen和acceptConn函数，外部包调用
// 开始监听本地的IP是否有连接，开启一个线程一直接收连接
func (connManager *ConnManager) Start(localIP string) {
	connManager.startListen(localIP)
	go connManager.acceptConn()
}

// startListen 开始监听某个端口
// 如果成功启动监听，返回true，否则返回false
func (connManager *ConnManager) startListen(localIP string) bool {
	connManager.LocalIP = localIP
	listen, err := net.Listen("tcp", connManager.LocalIP)
	if err != nil {
		log.Printf("开始监听时出现错误错误：%v", err)
		//TODO:错误处理
		return false
	}
	connManager.Listen = listen
	log.Println("监听成功")
	return true
}

// RequestConn 开始发起一个连接，如果连接发送不成功,就尝试每秒发送一次，尝试五次
// 五次连接不成功系统返回错误，连接成功新建peer，并新加入一个serverpeer
func (connManager *ConnManager) requestConn(targetIP string) *ServerPeer {
	serverPeer, ok := connManager.ServerPeers[targetIP]
	if ok {
		return serverPeer
	}
	var conn net.Conn
	var err error
	for i := 0; i < 5; i++ {
		conn, err = net.DialTimeout("tcp", targetIP, time.Second)
		if err != nil {
			log.Printf("建立连接不成功：%v", err)
			//TODO:错误处理
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

// acceptConn 接收一个连接,使用无限for循环，一直进行接收，需要开启另外的线程处理
func (connManager *ConnManager) acceptConn() {
	for {
		log.Println("正在Accept")
		conn, err := connManager.Listen.Accept()
		if err != nil {
			log.Printf("接收连接时发生错误：%v", err)
			//TODO:错误处理
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

// BroadCast 向所有建立的连接广播消息，外部包调用
func (connManager *ConnManager) BroadCast(msg *Message) {
	for _, serverPeer := range connManager.ServerPeers {
		serverPeer.SendMessage(msg)
	}
}

//Send 封装了请求连接和连接成功之后的发送，外部包调用
func (connManager *ConnManager) Send(msg *Message, targetIP string) {
	serverPeer := connManager.requestConn(targetIP)
	serverPeer.SendMessage(msg)
}
