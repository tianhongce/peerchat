package network

import (
	"log"
	"net"
	"time"
)

// ConnManager 保存map映射
type ConnManager struct {
	Listen     net.Listener
	LocalIP    string
	ServerPeer *ServerPeer
}

// Start 封装了startListen和acceptConn函数
func (connManager *ConnManager) Start(localIP string) {
	connManager.startListen(localIP)
	connManager.acceptConn()
}

// startListen 开始监听某个端口
func (connManager *ConnManager) startListen(localIP string) bool {
	connManager.LocalIP = localIP
	connManager.ServerPeer.peers = make(map[string]*Peer)
	listen, err := net.Listen("tcp", connManager.LocalIP)
	if err != nil {
		log.Printf("开始监听时出现错误错误：%v", err)
		return false
	}
	connManager.Listen = listen
	return true
}

// RequestConn 开始发起一个连接，如果连接发送不成功就for循环一直发，直到发送成功break出循环
func (connManager *ConnManager) RequestConn(targetIP string) {
	peer, ok := connManager.ServerPeer.peers[targetIP]
	if ok {
		connManager.ServerPeer.peer = peer //将此连接设置为当前peer
		return
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
	newPeer := Peer{ip, conn}
	log.Printf("创建一个peer, peer ip:%s", newPeer.targetIP)
	connManager.ServerPeer.peers[ip] = &newPeer //将新建连接加入到map中
	connManager.ServerPeer.peer = &newPeer      //将新建连接设置为当前peer
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
		peer := Peer{ip, conn}
		log.Printf("创建一个peer, peer ip:%s", peer.targetIP)
		go peer.RecvMsg()
		connManager.ServerPeer.peers[ip] = &peer
		connManager.ServerPeer.peer = &peer
	}
}
