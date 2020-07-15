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
	fmt.Println(compare(args[1], args[2]))
}

func compare(a, b string) bool {

	mapA := buildMap(a)
	mapB := buildMap(b)

	if len(mapA) != len(mapB) {
		return false
	}

	for k, v := range mapA {
		if v != mapB[k] {
			return false
		}
	}
	return true
}

func buildMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	return m
}
