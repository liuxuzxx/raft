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
	Id     string   `json:"id"`
	Domain string   `json:"domain"`
	Port   int      `json:"port"`
	Name   string   `json:"name"`
	Nodes  []Node   `json:"nodes"`
	Type   NodeType `json:"type"`
	Term   int      `json:"term"`
}

type Node struct {
	Id   string `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Term int    `json:"term"`
}

//
// Node的三种状态的类型
//
type NodeType int32

const (
	Follower  NodeType = 1
	Candidate NodeType = 2
	Leader    NodeType = 3
)

var Conf Config

func InitConfig(configPath string) {
	fmt.Println("开始加载配置信息....")
	viper.SetConfigFile(configPath)
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
	Conf.Server.Type = Follower
	fmt.Printf("查看加载的config配置信息:%s\n", string(jsonString))
}
