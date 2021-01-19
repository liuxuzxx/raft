package rest

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"raft/config"
	"raft/entity"
	"raft/server"
)

//
// election leader的API接口
//

func Vote(ctx iris.Context) {
	vote := &entity.VoteRequest{}
	if err := ctx.ReadJSON(vote); err != nil {
		_, _ = ctx.JSON(entity.VoteResponse{
			VoteId:   config.Conf.Server.Id,
			Result:   entity.Oppose,
			Describe: "解析Vote请求对象失败!",
		})
		return
	}
	_, _ = ctx.JSON(server.Election.ExecuteVote(*vote))
}

func MaintainAuthority(ctx iris.Context) {
	authority := &entity.LeaderAuthorityRequest{}
	if err := ctx.ReadJSON(authority); err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, _ = ctx.JSON(server.Election.FollowerLeader(*authority))
}
