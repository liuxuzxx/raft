package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"raft/server"
)

func main() {
	fmt.Println("Raft服务启动")
	app := route()
	_ = app.Listen(":12600")
}

func route() (app *iris.Application) {
	app = iris.New()

	v1 := app.Party("/v1.0.0/raft", corsConfig()).AllowMethods(iris.MethodOptions)
	{
		v1.Get("/information", server.RaftInformation)
		v1.PartyFunc("/election-leader", func(leaderParty router.Party) {
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
