package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	doc, err := findLinks("https://code.visualstudio.com/docs/editor/debugging")
	if err != nil {
		log.Printf("findlinks: %s", err)
	}

	forEachNode(doc, startElement, endElement)
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

func hasElementNodeChild(n *html.Node) bool {
	return n.FirstChild != nil || n.FirstChild.Type != html.ElementNode
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if hasElementNodeChild(n) {
			fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
			return
		}

		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
		return
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if hasElementNodeChild(n) {
			return
		}

		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		return
	}
}

func findLinks(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return doc, nil
}
