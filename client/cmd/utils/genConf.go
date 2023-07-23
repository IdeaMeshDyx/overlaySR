/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

type Config struct {
	Client struct {
		Address string `yaml:"address"`
		Ports   []int  `yaml:"ports"`
	} `yaml:"client"`
	Server struct {
		Address struct {
			IP   string `yaml:"ip"`
			Port int    `yaml:"port"`
		} `yaml:"address"`
	} `yaml:"server"`
}

const ConfigFileName = "client.yaml" // 配置文件名，我没找到从 viper 获取配置文件名的方法

var forcedFileOverwrite = false // 用来存放变量值

// genConfCmd represents the genConf command
var genConfCmd = &cobra.Command{
	Use:   "genConf",
	Short: "Create example configuration for people who do not know the format",
	Long:  `i do nto want to write too much.`,
	Run:   genYaml(),
}

func genYaml() {
	cfg := Config{
		Client: struct {
			Address string `yaml:"address"`
			Ports   []int  `yaml:"ports"`
		}{
			Address: "localhost",
			Ports:   []int{6023, 6029},
		},
		Server: struct {
			Address struct {
				IP   string `yaml:"ip"`
				Port int    `yaml:"port"`
			} `yaml:"address"`
		}{
			Address: struct {
				IP   string `yaml:"ip"`
				Port int    `yaml:"port"`
			}{
				IP:   "127.0.0.1",
				Port: 3099,
			},
		},
	}

	// 将Config实例序列化为YAML格式的字节数据
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		fmt.Printf("Error marshaling YAML data: %v\n", err)
		return
	}

	// 将YAML数据写入文件config.yaml
	err = ioutil.WriteFile("./config/config.yaml", data, 0644)
	if err != nil {
		fmt.Printf("Error writing YAML data to file: %v\n", err)
		return
	}

	fmt.Println("YAML data has been written to config.yaml")
}

func init() {
	rootCmd.AddCommand(genConfCmd)

	genConfCmd.Flags().BoolVarP(&forcedFileOverwrite, "forced", "f", false, "forced file overwrite") // 添加 -f 参数

}
