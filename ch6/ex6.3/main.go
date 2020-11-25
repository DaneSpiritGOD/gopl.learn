package main

import (
	"fmt"

	"github.com/DaneSpiritGOD/ch6/intset"
)

func main() {
	var x, y intset.IntSet

	x.AddAll(1, 144, 922, 188, 957, 1000)
	y.AddAll(77, 88, 99, 100, 922, 188, 1, 144)

	fmt.Printf("x: %s len: %d\n", x.String(), x.Len())
	fmt.Printf("y: %s len: %d\n", y.String(), y.Len())

	z := x.IntersectWith(&y)
	fmt.Printf("x IntersectWith y: %s len: %d\n", z.String(), z.Len())

	z = x.DifferenceWith(&y)
	fmt.Printf("x DifferenceWith y: %s len: %d\n", z.String(), z.Len())

	zsxy := x.SymmetricDifference(&y)
	fmt.Printf("x SymmetricDifference y: %s len: %d\n", zsxy.String(), zsxy.Len())

	zsyx := y.SymmetricDifference(&x)
	fmt.Printf("y SymmetricDifference x: %s len: %d\n", zsyx.String(), zsyx.Len())

	fmt.Printf("zsxy equals to zsyx: %v\n", zsxy.Equals(&zsyx))
}
