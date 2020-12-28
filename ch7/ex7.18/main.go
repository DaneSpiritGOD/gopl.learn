package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	url := "https://www.cnblogs.com"

	dec := newDecoder(fetchURL(url))

	var stack []string // stack of element names
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
			stack = append(stack, tok.Name.Local) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:

		}
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
