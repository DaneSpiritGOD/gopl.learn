package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Debug for string error
type errString struct {
	s string
}

func (s *errString) Error() string {
	return s.s
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	file := path.Base(resp.Request.URL.Path)
	if file == "/" {
		file = "index.html"
	}
	f, err := os.Create(file)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err != nil {
			err = closeErr
		} else {
			// set flat to true to test
			debugFlag := false
			if debugFlag {
				err = &errString{"This is fake err to test whether this defer function works."}
			}
		}
	}()

	return file, n, err
}

func main() {
	url := "https://books.studygolang.com/gopl-zh/ch5/ch5-08.html"
	local, n, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %s, error: %v\n", url, err)
		return
	}

	fmt.Printf("fetch %s -> %s, bytes: %d\n", url, local, n)
}
