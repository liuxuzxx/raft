package entity

//
// log replication日志复制这一节
// 是Leader被选出来之后，服务于客户端的操作
//
type CommandType string

const (
	Get CommandType = "get"
	Set CommandType = "set"
)
