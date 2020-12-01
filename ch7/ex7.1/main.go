package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// StringCounter alias as string
type StringCounter struct {
	WordCount int
	LineCount int
}

// GetWordCount return the number of wods
func GetWordCount(p []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return count
}

// GetLineCount return the number of line
func GetLineCount(p []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return count
}

func (sc *StringCounter) Write(p []byte) (n int, err error) {
	wordCount := GetWordCount(p)
	lineCount := GetLineCount(p)

	sc.WordCount += wordCount
	sc.LineCount += lineCount

	return len(p), nil
}

func (sc *StringCounter) String() string {
	return fmt.Sprintf("wordCount: %d, lineCount: %d", sc.WordCount, sc.LineCount)
}

func main() {
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	sc := &StringCounter{}
	fmt.Fprintf(sc, input)
	fmt.Println(sc)
}
