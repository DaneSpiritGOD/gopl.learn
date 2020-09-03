package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	SHA256 = "256"
	SHA384 = "384"
	SHA512 = "512"
)

var s = flag.String("s", "", "value to calc")
var algo = flag.String("algo", SHA256, "sha alogrithm")

func main() {
	flag.Parse()

	switch *algo {
	case SHA256:
		fmt.Printf("%x\n", sha256.Sum256([]byte(*s)))
	case SHA384:
		fmt.Printf("%x\n", sha512.Sum384([]byte(*s)))
	case SHA512:
		fmt.Printf("%x\n", sha512.Sum512([]byte(*s)))
	}
}
