package main

import "fmt"

func mergeSpace(source []byte) []byte {
	tailIndex := 0
	for index := 1; index < len(source); index++ {
		if source[index] != source[tailIndex] || !isSpace(source[index]) {
			tailIndex++
			if tailIndex != index {
				source[tailIndex] = source[index]
			}
		}
	}
	return source[:tailIndex+1]
}

func mergeSpaceS(source string) string {
	list := []byte(source)
	list = mergeSpace(list)
	return string(list)
}

func isSpace(b byte) bool {
	// This property isn't the same as Z; special-case it.
	switch b {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

func main() {
	fmt.Println(mergeSpaceS("aa bb"))
	fmt.Println(mergeSpaceS("aa   bb"))
	fmt.Println(mergeSpaceS("aa   bb c"))
	fmt.Println(mergeSpaceS("aa   bb c    d"))
	fmt.Println(mergeSpaceS("aa           d"))
	fmt.Println(mergeSpaceS("           d"))
}
