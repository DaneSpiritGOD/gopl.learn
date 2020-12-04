package main

import (
	"flag"
	"fmt"
	"github.com/DaneSpiritGOD/ex7.6/temperature"
)

var temp = temperature.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
