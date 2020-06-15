package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[0:], " ")
	fmt.Println(str)
}
