package network

// ServerPeer 应用层逻辑
// peer 当前peer连接
type ServerPeer struct {
	peer  *Peer
	peers map[string]*Peer
}

// SendMessage 将Message转码成buf并调用peer的发送来发送消息
func (serverPeer *ServerPeer) SendMessage(message *Message) {
	msgBuf := MsgToRlpData(*message)
	serverPeer.peer.SendMsg(msgBuf)
}

// ReceiveMessage 接收当前peer收到的buf数据并转码
func (serverPeer *ServerPeer) ReceiveMessage() {
	for {
		bufMsg := serverPeer.peer.RecvMsg()
		message := DataToRlpMsg(bufMsg)
		message.MsgFormat()
	}
}

// Broadcast 向所有已经建立的连接广播消息
func (serverPeer *ServerPeer) Broadcast(message *Message) {
	for _, peer := range serverPeer.peers {
		msgBuf := MsgToRlpData(*message)
		peer.SendMsg(msgBuf)
	}
}
