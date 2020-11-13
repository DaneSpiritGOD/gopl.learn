package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func max(vals ...int) int {
	if len(vals) == 0 {
		panic("no elements")
	}

	maxV := math.MinInt32
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}

	return maxV
}

func min(vals ...int) int {
	if len(vals) == 0 {
		panic("no elements")
	}

	minV := math.MaxInt32
	for _, v := range vals {
		if v < minV {
			minV = v
		}
	}

	return minV
}

func main() {
	conv := func() []int {
		var items []int
		for _, s := range os.Args[1:] {
			if i, err := strconv.Atoi(s); err == nil {
				items = append(items, i)
			}
		}
		return items
	}

	seq := conv()

	m := max(seq...)
	fmt.Printf("max number: %v\n", m)

	m = min(seq...)
	fmt.Printf("min number: %v\n", m)
}
