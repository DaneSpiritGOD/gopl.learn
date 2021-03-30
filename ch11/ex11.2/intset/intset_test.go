package intset_test

import (
	"testing"

	"github.com/DaneSpiritGOD/ex11.2/intset"
)

func TestAdd(t *testing.T) {
	var x intset.IntSet
	x.Add(1)
}

// func main( {
// 	var x, y intset.Intet
// 	x.Add1)
// 	x.Add(14)
// 	x.Add9)
// 	fmt.Printf("x: %s len: %d\n", x.String(), x.Len()) // "{1 9 14}"

// 	xCopy := x.Cop()
// 	fmt.Printf("xCopy: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 14}"

// 	xCopy.Remove(100)
// 	fmt.Printf("xCopy after remove 1000: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 14}"

// 	xCopy.Remove1)
// 	fmt.Printf("xCopy after remove 1: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 14}"

// 	xCopy.Clea()
// 	fmt.Printf("xCopy after clear: %s len: %d\n", xCopy.String(), xCopy.Len()) // "{1 9 14}

// 	y.Add9)
// 	y.Add(2)
// 	fmt.Printf("y: %s len: %d\n", y.String(), y.Len()) // "{1 9 14}"

// 	x.UnionWith(y)
// 	fmt.Printf("x after union with y: %s len: %d\n", x.String(), x.Len()) // "{1 9 42 14}"
// 	fmt.Println(x.Has(9), x.Has(123))                                     // "true fale"
// }
