package main

import "fmt"

func toNumber() (r int) {
	defer func() {
		if p := recover(); p != nil {
			r = 10
		}
	}()

	panic("return")
}

func main() {
	num := toNumber()
	fmt.Println(num)
}
