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

	l := len(s)

	for start, i := 0, (l-1)%3; i < l; start, i = i+1, i+3 {
		if start != 0 {
			buf.WriteByte(',')
		}

		buf.WriteString(s[start : i+1])
	}

	return buf.String()
}
