package main

import (
	"log"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	f, err := os.Create("trace.out")
	if err != nil {
		log.Println("create file failed:", err)
		return
	}
	defer f.Close()

	_ = trace.Start(f)
	defer trace.Stop()

	/*
		go func() {
			fabnic.Fab1(40)
			wg.Done()
		}()
		go func() {
			fabnic.Fab2(40)
			wg.Done()
		}()*/
	for i := 0; i < 10; i++ {
		go func(goid int) {
			var sum int
			for j := 0; j < 10000; j++ {
				sum += j
			}
			log.Println("go routines:", goid, "is finished")
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("all goroutines are finished")
}
