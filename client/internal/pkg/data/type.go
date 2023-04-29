package data

import (
	"bytes"
	"encoding/binary"
)

// Message 信息数据类型，用于定义在Client中解析并传输的数据类型
type Message interface {
	// Byte 转化为byte类型
	Byte() ([]byte, error)
}

type procMsg struct{}

type podMsg struct {
	podid        string
	podIndentity string
}

type nodeMsg struct{}

type serviceMsg struct{}

type netMsg struct{}

type WsMsg struct {
	id      string
	process procMsg
	pod     podMsg
	node    nodeMsg
	net     netMsg
}

func (ws *WsMsg) Byte() ([]byte, error) {
	// in: WsMsg. out: byte buffer
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, ws); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
