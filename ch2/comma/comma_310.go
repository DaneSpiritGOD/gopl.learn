package main

import (
	"bytes"
	"flag"
	"fmt"
)

var num = flag.String("n", "123", "number string")

func main() {
	flag.Parse()

	fmt.Println(comma(*num))
}

func comma(s string) string {
	var buf bytes.Buffer

	b := []byte(s)

	var start int
	if b[0] == '+' || b[1] == '-' {
		start = 1
	} else {
		start = 0
	}

	buf.Write(b[:start])

	dot := bytes.Index(b, []byte{'.'})
	if dot == -1 {
		commaCore(b[start:], &buf)
	} else {
		commaCore(b[start:dot], &buf)
		buf.Write(b[dot:])
	}

	return buf.String()
}

func commaCore(b []byte, buf *bytes.Buffer) {
	l := len(b)

	for start, i := 0, (l-1)%3; i < l; start, i = i+1, i+3 {
		if start != 0 {
			buf.WriteByte(',')
		}

		buf.Write(b[start : i+1])
	}
}
