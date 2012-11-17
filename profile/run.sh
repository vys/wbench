#!/bin/bash -x
pkill -9 redis-server
pkill -9 ab
pkill -9 wbench
 
redis-server > redis.log 2>&1 &
sleep 3
../wbench -profile web > wbench.log 2>&1 &
sleep 2
curl http://localhost:8081/reset > populate.log 2>&1
sleep 2
curl http://localhost:8081/test >> populate.log 2>&1
sleep 2
ab -c 100 -n 1000000 -r http://localhost:8081/ > ab.log 2>&1 
sleep 10
go tool pprof --pdf ../wbench web-cpu-1.prof > profile-cpu.pdf
go tool pprof --pdf ../wbench web-heap-1.prof > profile-heap.pdf
