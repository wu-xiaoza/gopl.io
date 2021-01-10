package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	//NewServeMux allocates and returns a new ServeMux.
	mux := http.NewServeMux()
	//Handle registers the handler for the given pattern. If a handler already exists for pattern, Handle panics.
	//The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	//ListenAndServe always returns a non-nil error.
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: $%d\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	//Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs. To check errors use ParseQuery.
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "$%d\n", price)
	} else {
		//WriteHeader sends an HTTP response header with the provided status code.
		//If WriteHeader is not called explicitly, the first call to Write will trigger an implicit WriteHeader(http.StatusOK). Thus explicit calls to WriteHeader are mainly used to send error codes.
		//The provided code must be a valid HTTP 1xx-5xx status code. Only one header may be written. Go does not currently support sending user-defined 1xx informational headers, with the exception of 100-continue response header that the Server sends automatically when the Request.Body is read.
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no sucn item: %q\n", item)
	}
}
