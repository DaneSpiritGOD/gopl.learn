package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("please type two strings")
		return
	}
	var a = 'c'
	fmt.Println(compare(args[1], args[2]))
}

func compare(a, b string) bool {
	return false
}
