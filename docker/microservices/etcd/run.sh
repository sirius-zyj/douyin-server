#!/bin/bash

etcd &

# 等待 etcd 启动
sleep 3

etcdctl put /config/config.yml "$(cat /config/config.yml)"

# 其他启动命令
# ...

# 等待 etcd 进程结束
wait