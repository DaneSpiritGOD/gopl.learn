package main

import (
	"fmt"

	intset "github.com/DaneSpiritGOD/ch6/IntSet"
)

func main() {
	var a intset.IntSet

	a.AddAll(1, 2, 3, 4, 100)

	fmt.Print("elements in a: ")
	for _, item := range a.Elems() {
		fmt.Printf("%d ", item)
	}
}
