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
