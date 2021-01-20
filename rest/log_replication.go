package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"raft/entity"
	"raft/server"
)

//
// log replication的Rest服务接口
//

func Command(ctx iris.Context) {
	command := &entity.CommandRequest{}
	if err := ctx.ReadJSON(command); err != nil {
		fmt.Println(err)
		panic(err)
	}
	server.LogReplication.Save(*command)
	_, _ = ctx.JSON("success save data!")
}
