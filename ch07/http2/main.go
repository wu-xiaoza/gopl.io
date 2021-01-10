package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	db := database{"shoes": 50, "socks": 5}
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	//The handler is typically nil, in which case the DefaultServeMux is used.
	//ListenAndServe always returns a non-nil error.
	log.Fatal(http.ListenAndServe("localhost:8080", db))
}

type dollars float32

//Sprintf formats according to a format specifier and returns the resulting string.
func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//URL specifies either the URI being requested (for server requests) or the URL to access (for client requests).
	//For server requests, the URL is parsed from the URI supplied on the Request-Line as stored in RequestURI. For most requests, fields other than Path and RawQuery will be empty. (See RFC 7230, Section 5.3)
	//For client requests, the URL's Host specifies the server to connect to, while the Request's Host field optionally specifies the Host header value to send in the HTTP request.
	switch req.URL.Path /* path (relative paths may omit leading slash) */	{
	case "/list":
		for item, price := range db {
			//Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		//Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs. To check errors use ParseQuery.
		//Get gets the first value associated with the given key. If there are no values associated with the key, Get returns the empty string. To access multiple values, use the map directly.
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			//WriteHeader sends an HTTP response header with the provided status code.
			//If WriteHeader is not called explicitly, the first call to Write will trigger an implicit WriteHeader(http.StatusOK). Thus explicit calls to WriteHeader are mainly used to send error codes.
			//The provided code must be a valid HTTP 1xx-5xx status code. Only one header may be written. Go does not currently support sending user-defined 1xx informational headers, with the exception of 100-continue response header that the Server sends automatically when the Request.Body is read.
			//HTTP status codes as registered with IANA. See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}