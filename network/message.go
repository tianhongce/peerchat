package network

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-ethereum/rlp"
)

// Message Message格式信息
// MsgText 信息文本
// SendTime 信息发送时间
// From 信息发送者
// To 信息接收者
type Message struct {
	MsgText  string
	SendTime string
	From     string
	To       string
}

// NewMessage 新建一个格式化的消息
func NewMessage(msgText string, sendTime string, from string, to string) *Message {
	newMessage := &Message{
		MsgText:  msgText,
		SendTime: sendTime,
		From:     from,
		To:       to,
	}
	return newMessage
}

// MsgFormat 将Message信息格式化输出
func (msg *Message) MsgFormat() {
	fmt.Printf("Received from: %s \n", msg.From)
	fmt.Printf("To: %s \n", msg.To)
	fmt.Printf("Time: %s \n", msg.SendTime)
	fmt.Printf("Text: %s \n", msg.MsgText)
}

// MsgToJSONData 将Message使用Json转码为[]byte
func MsgToJSONData(msg Message) []byte {
	msgData, err := json.Marshal(msg)
	if err != nil {
		log.Printf("MsgToJSONData时出现错误：%v", err)
	}
	return []byte(msgData)
}

// DataToJOSNMsg 将[]byte使用Json解码为Message
func DataToJOSNMsg(msgData []byte) Message {
	var msgText Message
	err := json.Unmarshal(msgData, &msgText)
	if err != nil {
		log.Printf("DataToJOSNMsg时出现错误：%v", err)
	}
	return msgText
}

//MsgToRlpData 将Message使用Rlp转码为[]byte
func MsgToRlpData(msg Message) []byte {
	msgData, err := rlp.EncodeToBytes(msg)
	if err != nil {
		log.Printf("MsgToRlpData时出现错误：%v", err)
	}
	return []byte(msgData)
}

// DataToRlpMsg 将[]byte使用Rlp解码为 Message
func DataToRlpMsg(msgData []byte) Message {
	var msgText Message
	err := rlp.DecodeBytes(msgData, &msgText)
	if err != nil {
		log.Printf("DataToRlpMsg时出现错误：%v", err)
	}
	return msgText
}
