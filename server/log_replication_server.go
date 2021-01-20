package server

import (
	"fmt"
	"raft/entity"
	"strconv"
)

//
// log replication的服务层
// log复制，持久化到本地
//
type LogReplicationServer struct {
	dataMap map[string]entity.DBEntry
	index   int
}

func (l *LogReplicationServer) Save(command entity.CommandRequest) {
	l.index = l.index + 1
	l.dataMap[command.Key+strconv.Itoa(l.index)] = entity.DBEntry{
		Index: l.index,
		Term:  0,
		Key:   command.Key,
		Value: command.Value,
	}
	l.debugDataMap()
}

func (l *LogReplicationServer) debugDataMap() {
	fmt.Println("----------------------------------------")
	for _, value := range l.dataMap {
		fmt.Printf("Index:%d,Term:%d,Key:%s,Value:%s\n", value.Index, value.Term, value.Key, value.Value)
	}
}

var LogReplication *LogReplicationServer

func init() {
	LogReplication = &LogReplicationServer{
		dataMap: make(map[string]entity.DBEntry, 0),
		index:   0,
	}
}
