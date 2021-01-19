package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"raft/config"
	"raft/entity"
	"strconv"
	"sync"
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
	Id         string
	Term       int
	Type       config.NodeType
	VoteId     string
	IsVote     bool
	AgreeCount int
	lock       sync.Mutex
}

func (e *ElectionLeader) ExecuteVote(vote entity.VoteRequest) entity.VoteResponse {
	e.lock.Lock()
	defer e.lock.Unlock()
	if e.IsVote || e.Type != config.Follower {
		return entity.VoteResponse{
			VoteId:   e.Id,
			Result:   entity.Oppose,
			Describe: "我不是Follower,或者是已经给别人投递过票了，所以不能给您投票了！",
		}
	}
	e.VoteId = vote.Id
	e.IsVote = true
	e.Type = config.Follower
	return entity.VoteResponse{
		VoteId:   e.Id,
		Result:   entity.Agree,
		Describe: "给您投票，我现在是您的Follower！",
	}
}

func (e *ElectionLeader) FollowerLeader(authority entity.LeaderAuthorityRequest) entity.FollowerAuthorityResponse {
	fmt.Printf("节点:%s 接收到了Leader的权威心跳:%s，任期为:%d\n", e.Id, authority.LeaderId, authority.Term)
	e.Type = config.Follower
	e.IsVote = true
	return entity.FollowerAuthorityResponse{
		FollowerId: e.Id,
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
	e.initiateVote()
}

func (e *ElectionLeader) initiateVote() {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.switchRole()
	e.sendVote()
	fmt.Printf("节点：%s 获取的投票数量是：%d\n", e.Id, e.AgreeCount)
	e.maintainAuthority()
}

func (e *ElectionLeader) sendVote() {
	for _, node := range config.Conf.Server.Nodes {
		client := &http.Client{}
		jsonBytes, _ := json.Marshal(entity.VoteRequest{
			Id:       e.Id,
			Term:     e.Term + 1,
			Describe: "请为我投票！",
		})
		url := "http://" + node.Ip + ":" + strconv.Itoa(node.Port) + "/v1.0.0/raft/election-leader/vote"
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
		if err != nil {
			fmt.Println(err)
		}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Printf("发向:%s 的投票出现错误，可能是node没有启动导致的！\n", node.Id)
			continue
		}
		responseBytes, _ := ioutil.ReadAll(resp.Body)
		voteResponse := &entity.VoteResponse{}
		_ = json.Unmarshal(responseBytes, voteResponse)
		if voteResponse.Result == entity.Agree {
			e.AgreeCount = e.AgreeCount + 1
			fmt.Printf("投票者：%s 投了我一票\n", voteResponse.VoteId)
		} else {
			fmt.Printf("投票者：%s 给我投递了反对票！\n", voteResponse.VoteId)
		}
	}
}

func (e *ElectionLeader) switchRole() {
	if !e.IsVote {
		e.Type = config.Candidate
		e.IsVote = true
		e.AgreeCount = 1
	}
}

func (e *ElectionLeader) maintainAuthority() {
	if e.AgreeCount > (len(config.Conf.Server.Nodes)+1)/2 {
		for _, node := range config.Conf.Server.Nodes {
			client := &http.Client{}
			jsonBytes, _ := json.Marshal(entity.LeaderAuthorityRequest{
				LeaderId: e.Id,
				Term:     e.Term,
			})
			url := "http://" + node.Ip + ":" + strconv.Itoa(node.Port) + "/v1.0.0/raft/election-leader/follower"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
			if err != nil {
				fmt.Println(err)
			}
			resp, err := client.Do(request)
			if err != nil {
				fmt.Printf("发向:%s 的Authority权威建立出现问题，可能是node没有启动导致的！\n", node.Id)
				continue
			}
			responseBytes, _ := ioutil.ReadAll(resp.Body)
			followerResponse := &entity.FollowerAuthorityResponse{}
			_ = json.Unmarshal(responseBytes, followerResponse)
			fmt.Printf("查看Follower的返回结果:%s\n", followerResponse.FollowerId)
		}
	}
}

func randomMillis() time.Duration {
	rand.Seed(time.Now().UnixNano())
	interval := rand.Intn(150) + 150
	fmt.Printf("获取的随机时间是:%d\n", interval)
	return time.Millisecond * time.Duration(interval)
}

var Election *ElectionLeader

func InitElectionLeader() {
	Election = &ElectionLeader{
		Id:     config.Conf.Server.Id,
		Term:   config.Conf.Server.Term,
		Type:   config.Conf.Server.Type,
		IsVote: false,
	}

	go Election.triggerElection()
}
