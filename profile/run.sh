#!/bin/bash -x
pkill -9 redis-server
pkill -9 ab
pkill -9 wbench
sleep 2
redis-server > redis.log 2>&1 &
sleep 2
../wbench -profile web > wbench.log 2>&1 &
sleep 2
curl http://localhost:8081/reset > populate.log 2>&1
sleep 2
curl http://localhost:8081/test >> populate.log 2>&1
sleep 2
ab -c 100 -n 1000000 -r http://localhost:8081/ > ab.log 2>&1 &
go tool pprof --pdf http://localhost:8081/debug/pprof/profile?seconds=300 > profile.pdf &
sleep 10
go tool pprof --pdf http://localhost:8081/debug/pprof/heap > heap-10.pdf
sleep 90
go tool pprof --pdf http://localhost:8081/debug/pprof/heap > heap-100.pdf
sleep 100
go tool pprof --pdf http://localhost:8081/debug/pprof/heap > heap-200.pdf
sleep 100
go tool pprof --pdf http://localhost:8081/debug/pprof/heap > heap-300.pdf
