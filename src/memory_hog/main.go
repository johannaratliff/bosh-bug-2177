package main

import (
	"log"
	"os"
	"time"
)

var SHA string

func main() {
	loggr := log.New(os.Stderr, "", log.LUTC)

	loggr.Println("Starting memory hog process...")
	loggr.Println(time.UnixDate)
	loggr.Printf("SHA: %s \n", SHA)

	for {
		loggr.Println("Starting eating up memory")
		time.Sleep(100 * time.Millisecond) // make sure the CPU isn't overloaded
		go func() {
			for {
				var slice []int64
				time.Sleep(100 * time.Millisecond) // make sure the CPU isn't overloaded
				slice = append(slice, time.Now().UnixNano())
			}
		}()
	}
}
