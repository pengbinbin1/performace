package main

import (
	"fabnic"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func main() {
	log.Println("start main thread...")
	f, err := os.Create("pprof.out")
	if err != nil {
		log.Println("create file failed:", err)
		return
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fabnic.Fab1(40)
		wg.Done()
	}()

	go func() {
		fabnic.Fab2(40)
		wg.Done()
	}()
	wg.Wait()
}
