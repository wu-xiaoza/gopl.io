package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//WaitForServer .
func WaitForServer(url string) error {
	// Common durations. There is no definition for units of Day or larger to avoid confusion across daylight savings time zone transitions.
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	// Before reports whether the time instant t is before u.
	for tries := 0; time.Now().Before(deadline); tries++ {
		/* Head issues a HEAD to the specified URL. If the response is one of the following redirect codes, 
		Head follows the redirect, up to a maximum of 10 redirects: */
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		// Sleep pauses the current goroutine for at least the duration d. A negative or zero duration causes Sleep to return immediately.
		// Common durations. There is no definition for units of Day or larger to avoid confusion across daylight savings time zone transitions.
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}

/* go run . https://golang.org
2020/12/24 12:17:53 server not responding (Head "https://golang.org": dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.); retrying...
2020/12/24 12:18:15 server not responding (Head "https://golang.org": dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.); retrying...
2020/12/24 12:18:38 server not responding (Head "https://golang.org": dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.); retrying...
Site is down: server https://golang.org failed to respond after 1m0s
exit status 1 */
