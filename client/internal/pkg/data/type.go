package data

import (
	"encoding/json"
	"log"
)

// Message 信息数据类型，用于定义在Client中解析并传输的数据类型
type Message interface {
	// Byte 转化为byte类型
	Byte() ([]byte, error)
}

type procMsg struct{}

type podMsg struct {
	PodId        string `json:"podId"`
	PodIndentity string `json:"podIndentity"`
}

type nodeMsg struct{}

type serviceMsg struct{}

type netMsg struct{}

type WsMsg struct {
	Id      string  `json:"id" :"id"`
	Process procMsg `json:"process" :"process"`
	Pod     podMsg  `json:"pod" :"pod"`
	Node    nodeMsg `json:"node"`
	Net     netMsg  `json:"net"`
}

func (ws WsMsg) Byte() ([]byte, error) {
	// in: WsMsg. out: byte buffer
	msg, err := json.Marshal(ws)
	if err != nil {
		log.Fatalf("ws To Josn Failed")
	}
	return msg, nil
}
