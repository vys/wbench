package main

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"errors"
	"log"
	"math"
	"redis"
)

var loadtimeout uint64 = 0
var savetimeout uint64 = 0

func load(r redis.AsyncClient, k string) (obj interface{}, err error) {
	err = errors.New("RedisTimeout")
	val, timeout := RedisGet(r, k)
	if timeout {
		return
	}
	zr, err := zlib.NewReader(bytes.NewReader(val))
	if err != nil {
		log.Fatal("Failed to create zlib reader with error: ", err)
	}
	defer zr.Close()
	jd := json.NewDecoder(zr)
	err = jd.Decode(&obj)
	if err != nil {
		log.Fatal("Failed to decode json with error: ", err)
	}
	return
}

func save(r redis.AsyncClient, k string, obj interface{}) {
	var b bytes.Buffer
	z := zlib.NewWriter(&b)
	defer z.Close()
	je := json.NewEncoder(z)
	err := je.Encode(obj)
	if err != nil {
		log.Fatal("Failed to json Encode with error: ", err)
	}
	z.Flush()
	RedisSet(r, k, b.Bytes())
}

func compute() {
	var k float64
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			k += math.Sin(float64(i)) * math.Sin(float64(j))
		}
	}
	_ = k
}
