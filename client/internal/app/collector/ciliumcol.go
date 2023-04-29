package collector

import "overlaysr/client/internal/pkg/data"

type Collector struct {
	/* TODO:
	* 实现Collector 接口
	 */
}

type Message interface {
	Collect(buffer chan data.Message)
}

func (c *Collector) Collect(buffer chan data.Message) {

}
