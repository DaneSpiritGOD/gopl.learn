package main

import (
	"fmt"

	"github.com/DaneSpiritGOD/ch6/intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Printf("x: %s len: %d\n", x.String(), x.Len()) // "{1 9 144}"

	xCopy := x.Copy()
	fmt.Printf("xCopy: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 144}"

	xCopy.Remove(1000)
	fmt.Printf("xCopy after remove 1000: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 144}"

	xCopy.Remove(1)
	fmt.Printf("xCopy after remove 1: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 144}"

	xCopy.Clear()
	fmt.Printf("xCopy after clear: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Printf("y: %s len: %d\n", y.String(), y.Len()) // "{1 9 144}"

	x.UnionWith(&y)
	fmt.Printf("x after union with y: %s len: %d\n", x.String(), x.Len()) // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123))                                     // "true false"
}
