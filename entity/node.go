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
