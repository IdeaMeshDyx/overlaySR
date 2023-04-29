package data

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

// Message 信息数据类型，用于定义在Client中解析并传输的数据类型
type Message interface {
	// Byte 转化为byte类型
	Byte() ([]byte, error)
}

type proMsg struct{}

type podMsg struct {
	podId        string `json:"podId"`
	podIndentity string `json:"podIndentity"`
}

type nodeMsg struct{}

type serviceMsg struct{}

type netMsg struct{}

type WsMsg struct {
	id      string  `json:"id" :"id"`
	process proMsg  `json:"process" :"process"`
	pod     podMsg  `json:"pod" :"pod"`
	node    nodeMsg `json:"node"`
	net     netMsg  `json:"net"`
}

func (ws *WsMsg) Byte() ([]byte, error) {
	buf := new(bytes.Buffer)
	msg, err := json.Marshal(ws)
	if err != nil {
		log.Fatalf("ws To Josn Failed")
	}
	if err := binary.Write(buf, binary.BigEndian, msg); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
