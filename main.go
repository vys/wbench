package main

import (
	"fmt"
	"net/http"
	"runtime"
    _ "net/http/pprof"
)

const MAXPROCS = 16
const MAXCLIENTS = 10
const key = "user_data"

var pool *ClientPool

func resetHandler(w http.ResponseWriter, r *http.Request) {
	var obj interface{}
	obj = NewUser()
	client := pool.Get()
	defer pool.Release(client)
	save(client, key, obj)
	fmt.Fprintf(w, "OK! KEY %s VALUE %v", key, obj)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	client := pool.Get()
	defer pool.Release(client)

	var obj interface{}
	obj, _ = load(client, key)
	fmt.Fprintf(w, "OK! KEY %s VALUE %v", key, obj)
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	client := pool.Get()
	defer pool.Release(client)

	var obj interface{}
	obj, err := load(client, key)
	if err != nil {
		return
	}

	compute()

	save(client, key, obj)
	fmt.Fprintf(w, "OK! %s", key)
}

func main() {
	runtime.GOMAXPROCS(MAXPROCS)
	profile()

	pool = NewClientPool(MAXCLIENTS)

	http.HandleFunc("/", responseHandler)
	http.HandleFunc("/reset", resetHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/stats", statsHandler)

	http.ListenAndServe(":8081", nil)
}
