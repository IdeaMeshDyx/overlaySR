/*
package data

# This file defines the Message type

Author: DYX, ZJX

Date: 2023/07/22
*/
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

type proMsg struct{}

type podMsg struct{}

type nodeMsg struct{}

type serviceMsg struct{}

type netMsg struct{}

type WsMsg struct {
	id      string
	process proMsg
	pod     podMsg
	node    nodeMsg
	net     netMsg
}

func (ws *WsMsg) Byte() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, ws); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
