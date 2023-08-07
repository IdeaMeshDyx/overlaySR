/*
Copyright © 2023 YiXinDa 2374087322@qq.com  && JianXin Zhang
*/
package main

import (
	"flag"
	"overlaysr/client/cmd"

	klog "k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(flag.CommandLine)
	flag.Parse()
	defer klog.Flush()
	cmd.Execute()
}
