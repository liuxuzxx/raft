package entity

//
// election leader的实体类信息
//

type Vote struct {
	Id       string `json:"id"`
	Term     int    `json:"term"`
	Describe string `json:"describe"`
}

type VoteResult int

const (
	Agree = 1
	Oppose
)

type VoteResponse struct {
	VoteId   string     `json:"voteId"`
	Result   VoteResult `json:"result"`
	Describe string     `json:"describe"`
}
