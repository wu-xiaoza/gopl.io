package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling { //field
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	/* Header maps header keys to values. If the response had multiple headers with the same key, 
	they may be concatenated, with comma delimiters. (RFC 7230, section 3.2.2 requires that 
	multiple headers be semantically equivalent to a comma-delimited sequence.) When Header values 
	are duplicated by other fields in this struct (e.g., ContentLength, TransferEncoding, Trailer), 
	the field values are authoritative. */


	ct := resp.Header.Get("Content-Type")
	// HasPrefix tests whether the string s begins with prefix.
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	/* It implements the HTML5 parsing algorithm (https://html.spec.whatwg.org/multipage/syntax.html#tree-construction), 
	which is very complicated. The resultant tree can contain implicitly created nodes that have no explicit <tag> 
	listed in r's data, and nodes' parents can differ from the nesting implied by a naive processing of start and end 
	<tag>s. Conversely, explicit <tag>s in r's data can be silently dropped, with no corresponding node in the resulting tree. */


	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}

func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

// The Go Programming Language