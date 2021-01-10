package main

import (
	"log"
	"time")


func bigSlowOperation() {
	defer trace("bigSlowOperation")()

	//Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	//Now returns the current local time.
	start := time.Now()
	log.Printf("enter %s", msg)
	//Since returns the time elapsed since t. It is shorthand for time.Now().Sub(t).
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}