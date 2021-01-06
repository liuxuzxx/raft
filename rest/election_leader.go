package rest

import (
	"github.com/kataras/iris/v12"
	"raft/config"
	"raft/entity"
	"raft/server"
)

//
// election leader的API接口
//

func Vote(ctx iris.Context) {
	vote := &entity.Vote{}
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
