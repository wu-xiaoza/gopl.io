package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	/* It implements the HTML5 parsing algorithm (https://html.spec.whatwg.org/multipage/syntax.html#tree-construction), 
	which is very complicated. The resultant tree can contain implicitly created nodes that have no explicit <tag> 
	listed in r's data, and nodes' parents can differ from the nesting implied by a naive processing of start and end 
	<tag>s. Conversely, explicit <tag>s in r's data can be silently dropped, with no corresponding node in the resulting tree. */
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
