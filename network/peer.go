package network

import (
	"log"
	"net"
)

// Peer 每个已经建立了连接的端点
// peerIP 连接对方的IP地址
type Peer struct {
	targetIP string
	conn     net.Conn
}

// SendMsg 转码成为buf并发送
func (peer *Peer) SendMsg(byte []byte) {
	_, err := peer.conn.Write(byte)
	if err != nil {
		log.Printf("发送消息出现错误：%v", err)
	}
}

// RecvMsg 接收buf并解码成为message
func (peer *Peer) RecvMsg() []byte {
	buf := make([]byte, 1024)
	bufLan, err := peer.conn.Read(buf)
	var bufMsg []byte
	log.Println(buf)
	if err != nil {
		log.Printf("接收消息出现错误：%v", err)
	} else {
		bufMsg = (buf[0:bufLan])
	}
	return bufMsg
}
