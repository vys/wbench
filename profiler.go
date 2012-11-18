package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/vys/go-humanize"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

func profile() {
	profile := flag.String("profile", "", "write profile to file with following prefix")
	flag.Parse()
	if *profile != "" {
		go doprofile(*profile)
		//go GoRuntimeStats()
	}
}

type Stats map[string]interface{}

func GetRuntimeStats() Stats {
	m := new(runtime.MemStats)

	s := make(Stats)

	s["# goroutines"] = runtime.NumGoroutine()
	runtime.ReadMemStats(m)
	s["Memory Acquire"] = humanize.Bytes(m.Sys)
	s["Memory Used"] = humanize.Bytes(m.Alloc)
	s["# malloc"] = m.Mallocs
	s["# free"] = m.Frees
	s["GC enabled"] = m.EnableGC
	s["# GC"] = m.NumGC
	s["Last GC time"] = m.LastGC
	s["Next GC"] = humanize.Bytes(m.NextGC)

	return s
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	gc := r.FormValue("gc")
	if gc == "1" {
		runtime.GC()
	}
	jb, err := json.MarshalIndent(GetRuntimeStats(), " ", " ")

	if err != nil {
        fmt.Fprintf(w, "Error in json encoding of stats")
        return
	}
    w.Header().Set("Content-Type", "application/json")
    w.Write(jb)
}

func doprofile(fn string) {
	var err error
	var fc, fh, ft *os.File
	for i := 1; i > 0; i++ {
		fc, err = os.Create(fn + "-cpu-" + strconv.Itoa(i) + ".prof")
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(fc)
		time.Sleep(300 * time.Second)
		pprof.StopCPUProfile()
		fc.Close()

		fh, err = os.Create(fn + "-heap-" + strconv.Itoa(i) + ".prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(fh)
		fh.Close()

		ft, err = os.Create(fn + "-threadcreate-" + strconv.Itoa(i) + ".prof")
		if err != nil {
			log.Fatal(err)
		}
		pprof.Lookup("threadcreate").WriteTo(ft, 0)
		ft.Close()
		log.Println("Created CPU, heap and threadcreate profile of 300 seconds")
	}
}
