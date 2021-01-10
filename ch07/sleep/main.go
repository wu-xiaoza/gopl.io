package main

import (
	"flag"
	"fmt"
	"time"
)

//Duration defines a time.Duration flag with specified name, default value, and usage string. The return value is the address of a time.Duration variable that stores the value of the flag. The flag accepts a value acceptable to time.ParseDuration.
var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	//Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	// Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
	time.Sleep(*period)
	fmt.Println()
}