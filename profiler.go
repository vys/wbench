package main

import (
	"flag"
	"github.com/vys/go-humanize"
	"log"
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

func GoRuntimeStats() {
	m := new(runtime.MemStats)
	for {
		time.Sleep(5 * time.Second)
		log.Println("# goroutines: ", runtime.NumGoroutine())
		runtime.ReadMemStats(m)
		log.Println("Memory Acquired: ", humanize.Bytes(m.Sys))
		log.Println("Memory Used    : ", humanize.Bytes(m.Alloc))
		log.Println("# malloc       : ", m.Mallocs)
		log.Println("# free         : ", m.Frees)
		log.Println("GC enabled     : ", m.EnableGC)
		log.Println("# GC           : ", m.NumGC)
		log.Println("Last GC time   : ", m.LastGC)
		log.Println("Next GC        : ", humanize.Bytes(m.NextGC))
		runtime.GC()
	}
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
