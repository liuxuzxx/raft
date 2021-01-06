package server

import (
	"fmt"
	"time"
)

//
// election leader:leader的选举
//

func triggerElection() {
	duration := time.Second * 2
	timer := time.NewTimer(duration)
	defer timer.Stop()

	<-timer.C
	fmt.Println("开始election leader... start RPC vote")
}

func init() {
	go triggerElection()
}
