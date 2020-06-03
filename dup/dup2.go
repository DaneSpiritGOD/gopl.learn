package main

import (
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {

	}
}

func countLines(f *os.File, counts map[string]int) {
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		counts[buf.Text()]++
	}
}
