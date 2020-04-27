package main

import (
	"fabnic"
	"log"
	"time"
)

func main() {
	log.Println("Enter main thread")

	t := time.Now()
	res := fabnic.Fab1(40)
	log.Println("fab1 result:", res)
	log.Println("fab1 cost time:", time.Since(t))

	log.Println("call fab2")
	t2 := time.Now()
	res2 := fabnic.Fab2(40)
	log.Println("fab2 result:", res2)
	log.Println("fab2 cost time:", time.Since(t2))
}
