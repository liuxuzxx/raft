package entity

//
// 系统基本配置信息
//
type Raft struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Detail  string `json:"detail"`
}
