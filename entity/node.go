package entity

//
// 系统基本配置信息
//
type Raft struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Detail  string `json:"detail"`
}

//
// Node的三种状态的类型
//
type NodeType int32

const (
	Follower  NodeType = 1
	Candidate NodeType = 2
	Leader    NodeType = 3
)
