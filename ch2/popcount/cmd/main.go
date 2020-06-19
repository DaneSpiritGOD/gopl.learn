package main

import (
	"ch2/popcount"
	"fmt"
)

func main() {
	fmt.Printf("%d\n", popcount.PopCount1(897864789123))
	fmt.Printf("%d\n", popcount.PopCount2(897864789123))
}
