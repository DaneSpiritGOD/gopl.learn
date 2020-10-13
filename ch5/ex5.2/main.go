// Findlinks1 prints the links in an HTML document read from standard input.

//.\ch1\fetch\fetch1.exe https://code.visualstudio.com/docs/editor/debugging | .\ch5\ex5.2\main.exe
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	m := make(map[string]int)
	visit(m, doc)
	for a, count := range m {
		fmt.Printf("%v %v\n", a, count)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		links[n.Data]++
	}

	// for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 	links = visit(links, c)
	// }

	if n.FirstChild != nil {
		visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(links, n.NextSibling)
	}
}
