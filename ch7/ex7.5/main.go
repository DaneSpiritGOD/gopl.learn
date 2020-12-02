package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type limitedReader struct {
	R io.Reader
	N int64
}

func limitReader(r io.Reader, n int64) io.Reader {
	return &limitedReader{r, n}
}

func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}

	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {
	buf := new(bytes.Buffer)
	buf.ReadFrom(os.Stdin)

	r := limitReader(buf, 5)

	p := [1000]byte{}
	n, _ := r.Read(p[:])
	fmt.Printf("limit read from stdin: %s", string(p[0:n]))
}
