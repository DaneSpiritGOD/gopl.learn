package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//Head https prefix
const Head string = "https://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, Head) {
			url = Head + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetach: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("Status Code: %d\n", resp.StatusCode)
	}
}
