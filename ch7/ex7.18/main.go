package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type node interface{} // CharData or *Element

type charData string

type element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []node
}

func main() {
	url := "https://www.cnblogs.com"

	data := fetchURL(url)
	dec := newDecoder(data)
	root := buildRoot(dec)

	fmt.Println(root)
}

func buildRoot(dec *xml.Decoder) node {
	var root node

	var stack []*element // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			e := &element{
				Type: tok.Name,
				Attr: tok.Attr,
			}

			if root == nil {
				root = e
			} else {
				last := stack[len(stack)-1]
				last.Children = append(last.Children, e)
			}
			stack = append(stack, e) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) > 0 {
				s := strings.TrimSpace(string(tok))
				if s != "" {
					last := stack[len(stack)-1]
					last.Children = append(last.Children, charData(s))
				}
			}
		}
	}

	return root
}

func (n *element) String() string {
	b := &bytes.Buffer{}
	visit(n, b, 0)
	return b.String()
}

func visit(n node, w io.Writer, depth int) {
	switch n := n.(type) {
	case *element:
		fmt.Fprintf(w, "%*s%s %s\n", depth*2, "", n.Type.Local, n.Attr)
		for _, c := range n.Children {
			visit(c, w, depth+1)
		}
	case charData:
		fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func fetchURL(url string) []byte {
	resp, err := http.Get(url)
	defer func() {
		resp.Body.Close()
	}()

	if err != nil {
		panic(fmt.Sprintf("fetch: %v\n", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("fetch: reading %s: %v\n", url, err))
	}

	return body
}

func newDecoder(data []byte) *xml.Decoder {
	dec := xml.NewDecoder(bytes.NewBuffer(data))
	dec.Strict = false
	dec.AutoClose = xml.HTMLAutoClose
	dec.Entity = xml.HTMLEntity

	return dec
}
