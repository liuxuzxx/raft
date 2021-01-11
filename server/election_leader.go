package server

import (
	"fmt"
	"math/rand"
	"raft/config"
	"raft/entity"
	"time"
)

//
// 投票处理对象
// 我需要关注的事情是：
// 1.投票请求过来了，我需要查看自己是否已经投票了
// 2.如果已经投票过了，那么就返回反对票的操作
// 3.我给谁投了票，我需要记录一下，打印日志方便查找问题的所在
// 问题是：
// 1、投票之后，下一次什么时候生效，不能一直都是投票过后的状态
type ElectionLeader struct {
	Id     string
	Term   int
	Type   config.NodeType
	VoteId string
	IsVote bool
}

func (e *ElectionLeader) ExecuteVote(vote entity.VoteRequest) entity.VoteResponse {
	if e.IsVote {
		return entity.VoteResponse{
			VoteId:   e.Id,
			Result:   entity.Oppose,
			Describe: "已经给别人投递过票了，所以不能给您投票了！",
		}
	}
	e.VoteId = vote.Id
	e.IsVote = true
	return entity.VoteResponse{
		VoteId:   e.Id,
		Result:   entity.Agree,
		Describe: "给您投票，我现在是您的Follower！",
	}
}

//
// election leader:leader的选举
//

func (e *ElectionLeader) triggerElection() {
	timer := time.NewTimer(randomMillis())
	defer timer.Stop()
	<-timer.C
	fmt.Println("开始election leader... start RPC vote")
}

func randomMillis() time.Duration {
	rand.Seed(time.Now().UnixNano())
	interval := rand.Intn(150) + 150
	fmt.Printf("获取的随机时间是:%d\n", interval)
	return time.Millisecond * time.Duration(interval)
}

var Election ElectionLeader

func init() {
	Election = ElectionLeader{
		Id:     config.Conf.Server.Id,
		Term:   config.Conf.Server.Term,
		Type:   config.Conf.Server.Type,
		IsVote: false,
	}

	go Election.triggerElection()
}
