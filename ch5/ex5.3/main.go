// Findlinks1 prints the links in an HTML document read from standard input.

//..\..\ch1\fetch\fetch1.exe https://code.visualstudio.com/docs/editor/debugging | .\main.exe
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	visit(os.Stdout, doc)
}

// visit appends to links each link found in n and returns the result.
func visit(w io.Writer, n *html.Node) {
	if n.Type == html.TextNode {
		if n.Parent != nil && n.Parent.Data != "script" && n.Parent.Data != "style" {
			w.Write([]byte(n.Data))
		}
	}

	if n.FirstChild != nil {
		visit(w, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(w, n.NextSibling)
	}
}
