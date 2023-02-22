package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	baton := make(chan int)
	go runner(done, baton)
	fmt.Println("Relay race is starting...")
	baton <- 1
	<-done
	fmt.Println("Relay race is over...")
}

func runner(done chan<- bool, baton chan int) {
	runnum := <-baton
	fmt.Printf("Runner %d is starting to run...\n", runnum)
	if runnum < 4 {
		go runner(done, baton)
		fmt.Printf("\tRunner %d is at the starting line, waiting for the baton...\n", runnum+1)
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("Runner %d has finished running\n", runnum)
	if runnum < 4 {
		fmt.Printf("\tRunner %d is passing the baton to runner %d\n", runnum, runnum+1)
		baton <- runnum + 1
	} else {
		done <- true
	}
}
