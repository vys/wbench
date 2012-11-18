package main

import (
	"github.com/vys/Go-Redis/redis"
	"log"
	"time"
)

type ClientPool struct {
	size       int
	clientChan chan redis.AsyncClient
}

func (p *ClientPool) Get() redis.AsyncClient {
	return <-p.clientChan
}

func (p *ClientPool) Release(c redis.AsyncClient) {
	p.clientChan <- c
}

func NewClientPool(size int) *ClientPool {
	p := &ClientPool{size: size}
	p.clientChan = make(chan redis.AsyncClient, size)
	for i := 0; i < size; i++ {
		client := NewRedisAsyncClient()
		p.clientChan <- client
	}
	return p
}

func NewRedisAsyncClient() redis.AsyncClient {
	spec := redis.DefaultSpec().Db(0).Host("127.0.0.1")
	client, err := redis.NewAsynchClientWithSpec(spec)
	if err != nil {
		panic(err)
	}
	return client
}

const REDIS_GET_TIMEOUT = 20 * time.Millisecond
const REDIS_SET_TIMEOUT = 20 * time.Millisecond

func RedisGet(r redis.AsyncClient, k string) (value []byte, timeout bool) {
	f, rerr := r.Get(k)
	if rerr != nil {
		log.Fatal("RedisGet failed: ", rerr)
	}

	value, rerr, timeout = f.TryGet(REDIS_GET_TIMEOUT)
	if rerr != nil {
		log.Fatal("RedisGet failed: ", rerr)
	}

	if timeout {
		loadtimeout++
		log.Println("load timeout! count: ", loadtimeout)
		return
	}
	return
}

func RedisSet(r redis.AsyncClient, k string, val []byte) (timeout bool) {
	f, rerr := r.Set(key, val)
	if rerr != nil {
		log.Fatal("RedisSet failed: ", rerr)
	}

	_, rerr, timeout = f.TryGet(REDIS_SET_TIMEOUT)
	if rerr != nil {
		log.Fatal("RedisSet failed: ", rerr)
	}

	if timeout {
		savetimeout++
		log.Println("save timeout! count: ", savetimeout)
		return
	}

	return
}
