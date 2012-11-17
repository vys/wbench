#!/bin/bash -xe
redis-server > redis.log 2>&1 &
../wbench -profile web > wbench.log 2>&1 &
curl http://localhost:8081/reset > populate.log 2>&1
curl http://localhost:8081/test >> populate.log 2>&1
ab -c 100 -n 10000000 -r http://localhost:8081/
go tool pprof --pdf wbench wbench-cpu-1.prof > profile-cpu.pdf
go tool pprof --pdf wbench wbench-heap-1.prof > profile-heap.pdf
