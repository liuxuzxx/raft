package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"os"
	"raft/config"
	"raft/rest"
	"raft/server"
	"strconv"
)

func main() {
	fmt.Println("Raft服务启动")
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("请携带上config文件的地址参数信息...")
		return
	}
	fmt.Printf("传入的配置文件路径参数:%s\n", args[1])
	config.InitConfig(args[1])
	app := route()
	_ = app.Run(iris.Addr(config.Conf.Server.Domain + ":" + strconv.Itoa(config.Conf.Server.Port)))
}

func route() (app *iris.Application) {
	app = iris.New()

	v1 := app.Party("/v1.0.0/raft", corsConfig()).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/information", server.RaftInformation)
		v1.PartyFunc("/election-leader", func(leaderParty router.Party) {
			leaderParty.Post("/vote", rest.Vote)
		})
	}
	return app
}

func corsConfig() iris.Handler {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	return crs
}
