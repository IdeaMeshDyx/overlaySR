/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// 在 root.go 中就直接使用了 cobra 命令来初始化 rootCmd 结构，CLI 中的其他所有命令都将是 rootCmd 这个根命令的子命令。
var rootCmd = &cobra.Command{
	Use:   "oversr",
	Short: "oversr_client deployed among all the nodes ",
	Long:  `oversr_client deployed among all the nodes , it consists of three parts: CNI Part, WS Part,Control Part`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("overlay start to communicate with underlay")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// rootCmd 根命令就会首先运行 initConfig 函数，当所有的初始化函数执行完成后，才会执行 rootCmd 的 RUN: func 执行函数
func init() {
	cobra.OnInitialize(initConfig)

	fmt.Println("inside initConfig")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// 全局添加的参数
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// 局域添加的参数
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find current directory.
		home, err := os.Getwd()
		cobra.CheckErr(err)
		path := home + "/config/"
		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(path)
		fmt.Printf("the config is : %s", path)
		viper.SetConfigType("yaml")
		viper.SetConfigName("client")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
