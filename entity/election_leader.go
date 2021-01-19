package entity

//
// election leader的实体类信息
//

type VoteRequest struct {
	Id       string `json:"id"`
	Term     int    `json:"term"`
	Describe string `json:"describe"`
}

type VoteResult int

const (
	Agree  VoteResult = 1
	Oppose VoteResult = 2
)

type VoteResponse struct {
	VoteId   string     `json:"voteId"`
	Result   VoteResult `json:"result"`
	Describe string     `json:"describe"`
}

//
//当leader选举出来之后，需要周期的发送心跳给其他
//的没有在这个Term当选的node，来建立权威
//
type LeaderAuthority struct {
	LeaderId string
	Term     int
}
