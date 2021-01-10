package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	/* 	Body represents the response body.The response body is streamed on demand as the Body field is read. 
	If the network connection fails or the server terminates the response, Body.Read calls return an error.
	The http Client and Transport guarantee that Body is always non-nil, even on responses without a body or 
	responses with a zero-length body. It is the caller's responsibility to close Body. The default HTTP 
	client's Transport may not reuse HTTP/1.x "keep-alive" TCP connections if the Body is not read to completion 
	and closed.The Body is automatically dechunked if the server replied with a "chunked" Transfer-Encoding.
	As of Go 1.12, the Body will also implement io.Writer on a successful "101 Switching Protocols" response, 
	as used by WebSockets and HTTP/2's "h2c" mode.*/	
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}