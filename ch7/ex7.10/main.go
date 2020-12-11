package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(os.Args)

	argLen := len(os.Args)
	if argLen == 1 {
		panic("need input []int numbers")
	}

	var s []int = nil
	for i := 1; i < argLen; i++ {
		num, err := strconv.Atoi(os.Args[i])
		if err != nil {
			continue
		}
		s = append(s, num)
	}

	fmt.Printf("slice isPalindrome result: %t", isPalindrome(sort.IntSlice(s)))
}

func isPalindrome(s sort.Interface) bool {
	n := s.Len()
	for i, j := 0, n-1; i < j; {
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}

		i++
		j--
	}
	return true
}
