/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"overlaysr/client/cmd"
)

func main() {
	viper.SetConfigName("client")    // 配置文件名称，不要写扩展名，不然找不到文件。看样子又是个bug v1.11.0
	viper.SetConfigType("yaml")      // 配置文件类型（扩展名）。必须写，不然找不到文件
	viper.AddConfigPath("./config/") // 放置配置文件的目录，可以添加多个
	err := viper.ReadInConfig()
	if err != nil {
		if !errors.As(err, new(viper.ConfigFileNotFoundError)) { // 配置文件没找到时使用默认配置或命令行参数
			log.Fatal(err)
		}
	}
	cmd.Execute()
}
