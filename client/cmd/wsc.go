/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// wscCmd represents the wsc command
var wscCmd = &cobra.Command{
	Use:   "wsc",
	Short: "fffff",
	Long:  `ffffff`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[3])
	},
}

func wsClient(cmd *cobra.Command, args string) {
	// create a new websocket agent
	// cil :=
	fmt.Println("IP:")
	fmt.Println("ws client start")

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
