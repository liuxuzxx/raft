package server

import (
	"fmt"
	"math/rand"
	"time"
)

//
// election leader:leader的选举
//

func triggerElection() {
	timer := time.NewTimer(randomMillis())
	defer timer.Stop()
	<-timer.C
	fmt.Println("开始election leader... start RPC vote")
}

func randomMillis() time.Duration {
	rand.Seed(time.Now().UnixNano())
	interval := rand.Intn(150) + 150
	fmt.Printf("获取的随机时间是:%d\n", interval)
	return time.Millisecond * time.Duration(interval)
}

func init() {
	go triggerElection()
}
