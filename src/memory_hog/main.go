package main

import (
	"log"
	"os"
	"time"
)

var sha string

func main() {
	loggr := log.New(os.Stderr, "", log.LUTC)

	loggr.Println("Starting memory hog process...")
	loggr.Printf("SHA: %s \n", sha)

	go func() {
		for {
			var slice []int64
			time.Sleep(100 * time.Millisecond) // make sure the CPU isn't overloaded
			slice = append(slice, time.Now().UnixNano())
		}
	}()
}
