package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letterCount := make(map[rune]int)
	digitCount := make(map[rune]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "letterCount, digitCount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}

		if unicode.IsLetter(r) {
			letterCount[r]++
			continue
		}
		if unicode.IsDigit(r) {
			digitCount[r]++
			continue
		}
	}

	fmt.Println("letter\tcount")
	for l, c := range letterCount {
		fmt.Printf("%q\t%d\n", l, c)
	}

	fmt.Println("digit\tcount")
	for d, c := range digitCount {
		fmt.Printf("%q\t%d\n", d, c)
	}
}
