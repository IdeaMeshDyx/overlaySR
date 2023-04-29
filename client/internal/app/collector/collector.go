package collector

import (
	"fmt"
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
	CiliumColl(buffer chan data.Message)
}

func (c Collector) CiliumColl(buffer chan data.Message) {

	// cilium API, send a pkg per second
	for {
		c.Msg.Id = fmt.Sprintf("%d", 1)
		buffer <- c.Msg
		time.Sleep(1 * time.Second)
	}

}
