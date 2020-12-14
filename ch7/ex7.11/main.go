package main

import (
	"fmt"
	"net/http"
)


type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

var db = database{"shoes": 50, "socks": 5}

func main() {

	http.Handle()
}
