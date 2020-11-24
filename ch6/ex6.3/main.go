package main

import (
	"fmt"

	"github.com/DaneSpiritGOD/ch6/intset"
)

func main() {
	var x, y intset.IntSet

	x.AddAll(1, 144, 922, 188, 957, 1000)
	y.AddAll(77, 88, 99, 100, 922, 188, 1, 144)
	z := x.IntersectWith(&y)
	fmt.Printf("x IntersectWith y: %s len: %d\n", z.String(), z.Len()) // "{1 9 144}"
}
