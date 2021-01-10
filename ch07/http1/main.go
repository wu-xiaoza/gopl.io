package main

import (
	"fmt"
	"log"
	"net/http"
)


type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	//Fatal is equivalent to Print() followed by a call to os.Exit(1).
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	//ListenAndServe always returns a non-nil error.
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}