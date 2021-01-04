package server

import (
	"github.com/kataras/iris/v12"
	"raft/entity"
)

//
// Raft系统的信息
//

func RaftInformation(ctx iris.Context) {
	_, _ = ctx.JSON(entity.Raft{
		Name:    "Raft服务",
		Version: "V1.0.0",
		Detail:  "使用Go语言实现Raft协议",
	})
}
