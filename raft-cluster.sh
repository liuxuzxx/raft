#!/bin/bash

doStart() {
  echo "开始启动Raft集群服务........."
  nohup ./raft ./config/config-12309.yml >nohup-12309.out 2>&1 & echo $! > raft-12309.pid
  nohup ./raft ./config/config-12310.yml >nohup-12310.out 2>&1 & echo $! > raft-12310.pid
  nohup ./raft ./config/config-12311.yml >nohup-12311.out 2>&1 & echo $! > raft-12311.pid
  nohup ./raft ./config/config-12312.yml >nohup-12312.out 2>&1 & echo $! > raft-12312.pid
  nohup ./raft ./config/config-12313.yml >nohup-12313.out 2>&1 & echo $! > raft-12313.pid
  echo "Raft集群服务启动完成！"
}

doStop() {
  echo "开始关闭Raft集群服务......"
  kill -9 `cat raft-12309.pid`
  kill -9 `cat raft-12310.pid`
  kill -9 `cat raft-12311.pid`
  kill -9 `cat raft-12312.pid`
  kill -9 `cat raft-12313.pid`
  rm -rf nohup-*.out
  rm -rf raft-*.pid
  echo "Raft集群全都关闭了!"
}

case "$1" in
start)
  :
  doStart
  ;;
stop)
  :
  doStop
  ;;
*)
  :
  echo "命令的使用形式是：bash raft-cluster.sh start|stop!"
  ;;
esac

exit 0
