/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"overlaysr/client/internal/pkg/wsclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// wscCmd represents the wsc command
var wscCmd = &cobra.Command{
	Use:   "wsc",
	Short: "start WS Client",
	Long:  `start ws client ,this operation will start to communicate with server`,
	Run: func(cmd *cobra.Command, args []string) {
		req := wsclient.WsRequest{
			Ip:   viper.GetString("ip"),
			Port: viper.GetString("port"),
		}
		wsClient(req)
	},
}

func wsClient(req wsclient.WsRequest) {
	// create a new websocket agent
	// cil :=
	fmt.Printf("IP:%s, port: %s", req.Ip, req.Port)
	fmt.Printf("ws client start")

	//var hub agent.WsAgent
	// msg := hub.Message
	// fmt.Print(msg.Byte())
	// send a message using the agent
	//hub.Send()
}

func init() {
	rootCmd.AddCommand(wscCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wscCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wscCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
