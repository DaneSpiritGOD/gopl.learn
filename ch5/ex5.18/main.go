package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

type errString struct {
	s string
}

func (s *errString) Error() string {
	return s.s
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err *error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		return "", 0, &err1
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err1 := os.Create(local)
	if err1 != nil {
		return "", 0, &err1
	}
	n, err1 = io.Copy(f, resp.Body)

	defer func() {
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = &closeErr
		} else {
			err = &errString{"hello"}
		}
	}()

	return local, n, &err1
}

func main() {
	url := "https://books.studygolang.com/gopl-zh/ch5/ch5-08.html"
	local, n, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "featch: %s %v\n", url, err)
	}

	fmt.Printf("fetch %s -> %s, bytes: %d\n", url, local, n)
}
