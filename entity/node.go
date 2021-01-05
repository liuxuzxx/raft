package entity

//
// 服务节点的数据类型定义
//
type Node struct {
	Id    string `json:"id"`
	Ip    string `json:"ip"`
	Port  int    `json:"port"`
	Nodes []Node `json:"nodes"`
}

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
