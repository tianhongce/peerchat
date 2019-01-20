package network

import (
	"fmt"
	"log"
	"net"
)

// Peer 每个已经建立了连接的端点
// peerIP 连接对方的IP地址
type Peer struct {
	targetIP string
	conn     net.Conn
}

// NewPeer 新建一个Peer并根据参数初始化
func NewPeer(ip string, conn net.Conn) *Peer {
	peer := &Peer{
		targetIP: ip,
		conn:     conn,
	}
	return peer
}

// SendMsg 将buf格式的消息发送到连接peer，被ServerPeer调用
func (peer *Peer) SendMsg(byte []byte) {
	fmt.Println("正在发送消息")
	_, err := peer.conn.Write(byte)
	if err != nil {
		log.Printf("发送消息出现错误：%v", err)
		//TODO:错误处理
	}
}

// RecvMsg 接收连接的buf消息
func (peer *Peer) RecvMsg() []byte {
	buf := make([]byte, 1024)
	bufLan, err := peer.conn.Read(buf)
	var bufMsg []byte
	if err != nil {
		log.Printf("接收消息出现错误：%v", err)
		//TODO:错误处理
	} else {
		bufMsg = (buf[0:bufLan])
	}
	return bufMsg
}
