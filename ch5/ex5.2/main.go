// Findlinks1 prints the links in an HTML document read from standard input.

//..\..\ch1\fetch\fetch1.exe https://code.visualstudio.com/docs/editor/debugging | .\main.exe
package main

import (
	"fmt"
	"os"
	"sort"

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

	type kv struct {
		key   string
		value int
	}

	var pairs []kv
	for k, v := range m {
		pairs = append(pairs, kv{k, v})
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].value > pairs[j].value
	})

	for _, pair := range pairs {
		fmt.Printf("%v %v\n", pair.key, pair.value)
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
