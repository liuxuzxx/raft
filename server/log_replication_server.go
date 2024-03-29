package server

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"path/filepath"
	"raft/entity"
	log "raft/pb"
	"strconv"
	"sync"
)

//
// log replication的服务层
// log复制，持久化到本地
//
type LogReplicationServer struct {
	dataMap  map[string]log.LogEntry
	index    int64
	path     string
	fileName string
	mux      sync.RWMutex
}

func (l *LogReplicationServer) Save(command entity.CommandRequest) {
	l.mux.Lock()
	l.index = l.index + 1
	entry := &log.LogEntry{
		Index: l.index,
		Term:  0,
		Key:   command.Key,
		Value: command.Value,
	}
	l.dataMap[command.Key+strconv.FormatInt(l.index, 10)] = *entry
	l.appendLog(entry)
	l.mux.Unlock()
	//l.debugDataMap()
}

func (l *LogReplicationServer) debugDataMap() {
	fmt.Println("----------------------------------------")
	for _, value := range l.dataMap {
		fmt.Printf("Index:%d,Term:%d,Key:%s,Value:%s\n", value.Index, value.Term, value.Key, value.Value)
	}
}

func (l *LogReplicationServer) appendLog(logEntry *log.LogEntry) {
	logFile, err := os.OpenFile(filepath.Join(l.path, l.fileName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	write := bufio.NewWriter(logFile)
	jsonBytes, _ := proto.Marshal(logEntry)
	_, err = write.Write(jsonBytes)

	err = write.Flush()
	if err != nil {
		fmt.Println(err)
	}
}

func (l *LogReplicationServer) initLogReplication() {
	_ = os.MkdirAll(l.path, os.ModePerm)
	fmt.Printf("初始化log存储路径:%s\n", l.path)
}

var LogReplication *LogReplicationServer

func init() {
	LogReplication = &LogReplicationServer{
		dataMap:  make(map[string]log.LogEntry, 0),
		index:    0,
		path:     "./log",
		fileName: "log.raft",
	}
	LogReplication.initLogReplication()
}
