package data

import (
	"encoding/json"
	"time"
)

// Message 信息数据类型，用于定义在Client中解析并传输的数据类型
type Message interface {
	// Byte 转化为byte类型
	Byte() ([]byte, error)
}

type procMsg struct{}

type SinglePod struct {
	Id           int64    `json:"podId"`
	Indentity    int64    `json:"podIndentity"`
	Labels       []string `json:"labels"`
	LabelsSHA256 string   `json:"labelsSHA256"`
	// pod in cilium seems dont know hostaddr
	// HostAddr string   `json:"HostAddr"`
	UUID           []string `json:"UUID"`
	IPv4           string   `json:"IPv4"`
	IPv6           string   `json:"IPv6s"`
	Mac            string   `json:"Mac"`
	HostMac        string   `json:"HostMac"`
	InterfaceIndex int64    `json:"InterfaceIndex"`
	InterfaceName  string   `json:"InterfaceName"`
}
type PodsMsg struct {
	Pods []SinglePod `json:"pod"`
}
type nodeMsg struct{}

type serviceMsg struct{}

type netMsg struct{}

type WsMsg struct {
	Id      string     `json:"id"`
	Time    time.Time  `json:"time"`
	Process procMsg    `json:"process"`
	Pods    PodsMsg    `json:"pods"`
	Node    nodeMsg    `json:"node"`
	Svc     serviceMsg `json:"svc"`
	Net     netMsg     `json:"net"`
}

// convert websocket message to byte
func (ws WsMsg) Byte() ([]byte, error) {
	ws.Time = time.Now()
	msg, err := json.Marshal(ws)
	if err != nil {
		klog.Errorf("ws To Josn Failed: %v", err)
		return nil, err
	}
	return msg, err
}
