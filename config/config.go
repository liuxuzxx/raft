package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//
// 配置信息
//
type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	Domain string `json:"domain"`
	Port   int    `json:"port"`
	Name   string `json:"name"`
	Nodes  []Node `json:"nodes"`
}

type Node struct {
	Id   string `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Term int    `json:"term"`
}

var Conf Config

func init() {
	fmt.Println("开始加载配置信息....")
	viper.SetConfigFile("./config/config.yml")
	viper.SetConfigType("yml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Update the config file")
	})
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Parse the config file is error!")
	}
	err = viper.Unmarshal(&Conf)

	jsonString, _ := json.Marshal(&Conf)
	fmt.Printf("查看加载的config配置信息:%s", string(jsonString))
}
