/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"os/signal"
	"overlaysr/client/internal/pkg/wsclient"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the wsc command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start WS Client",
	Long:  `start ws client ,this operation will start to communicate with server`,
	Run: func(cmd *cobra.Command, args []string) {
		req := wsclient.WsRequest{
			Ip:   viper.GetString("ip"),
			Port: viper.GetString("port"),
		}
		if req.Ip == "" || req.Port == "" {
			klog.Error("get IP:Port failed")
		}
		klog.Infof("start WS Client at %v:%v", req.Ip, req.Port)
		startWS(req)
	},
}

func startWS(req wsclient.WsRequest) {
	// instantiate a collector and a websocket agent
	ws_buffer := make(chan data.Message, 10)
	defer close(ws_buffer)
	var agent agent.Agent
	agent.Init()
	coll := collector.Collector{
		CollID: "collector1",
		Msg:    data.WsMsg{},
		Exited: false,
	}

	// collector write data into this chan
	go coll.Collect(ws_buffer)

	go agent.Runing()
	// main func

	// agent read and send collector's data in this chan
	go agent.ReadAndSend(ws_buffer)

	done := make(chan struct{})
	defer close(done)
	interrupt := make(chan os.Signal, 1)
	defer signal.Stop(interrupt)
	signal.Notify(interrupt, os.Interrupt)
	// main thread to exit
	for {
		select {
		case <-done:
			log.Println("agent done")
			return
		case <-interrupt:
			log.Println("keyboard interrupt")
			agent.WsClient.Exited = true
			coll.Exited = true
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func init() {
	rootCmd.AddCommand(wscCmd)

}
