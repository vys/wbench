package main

import (
	"net/http"
	"runtime"
)

func statsHandler(w http.ResponseWriter, r *http.Request) {
	gc := r.FormValue("gc")
	if gc == "1" {
		runtime.GC()
	}
}
