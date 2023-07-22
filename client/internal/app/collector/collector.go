package collector

import (
	"overlaysr/client/internal/pkg/data"
	"time"
)

type Collector struct {
	/* TODO:
	* 实现Collector 接口
	 */
	CollID string
	Msg    data.WsMsg
}

type CollMsg interface {
	Collect(buffer chan data.Message)
}

func (c Collector) Collect(buffer chan data.Message) {

	// cilium API, send a pkg per second
	for {
		// c.Msg.Pods = ciliumAPI.GetEps()

		// for test, when in linux with cilium, delete below
		var pdata data.PodsMsg
		var psinge data.SinglePod
		psinge.Id = 1
		psinge.IPv4 = "127.1.1.1"
		psinge.InterfaceName = "for debug"
		pdata.Pods = append(pdata.Pods, psinge)
		c.Msg.Pods = pdata
		// for test, del above

		buffer <- c.Msg
		time.Sleep(1 * time.Second)
	}

}
