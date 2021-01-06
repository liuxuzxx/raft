package config

import (
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
}
