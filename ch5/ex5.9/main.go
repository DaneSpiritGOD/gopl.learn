package main

import (
	"fmt"
	"strings"
)

const key string = "foo"

func main() {
	f := func(s string) string {
		s += "_"
		return s
	}

	fmt.Println(expand("bbbfooaaafoccfoo", f))
}

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, key, f(key))
}
