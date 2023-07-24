package collector

import (
	"log"
	"overlaysr/client/internal/pkg/data"
	"time"
)

type Collector struct {
	CollID string
	Msg    data.WsMsg
	Exited bool
}

type CollMsg interface {
	Collect(buffer chan data.Message)
}

func (c Collector) Collect(buffer chan data.Message) {

	// cilium API, send a pkg per second
	for {
		// c.Msg.Pods = ciliumAPI.GetEps()
		if c.Exited {
			log.Printf("Collector: %s exited\n", c.CollID)
			break
		}
		// for test, when in linux with cilium, delete below
		var pdata data.PodsMsg
		var psinge data.SinglePod
		psinge.Id = 1
		psinge.IPv4 = "127.1.1.1"
		psinge.InterfaceName = "for debug"
		pdata.Pods = append(pdata.Pods, psinge)
		c.Msg.Pods = pdata
		// for test, del above

		log.Printf("Collector: %s: Returns Data\n", c.CollID)
		buffer <- c.Msg
		time.Sleep(1 * time.Second)
	}

}
