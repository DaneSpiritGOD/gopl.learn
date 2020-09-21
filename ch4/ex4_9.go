package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordCount[input.Text()]++
	}

	fmt.Println("word\tcount")
	for w, c := range wordCount {
		fmt.Printf("%q\t%d\n", w, c)
	}
}
