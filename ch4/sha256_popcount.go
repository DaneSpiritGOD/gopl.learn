package main

import (
	"crypto/sha256"
	"fmt"
)

func popCountByte(num byte) int {
	sum := 0
	for num != 0 {
		num = num & (num - 1)
		sum++
	}
	return sum
}

func compare(num1 byte, num2 byte) int {
	if num1 == num2 {
		return 0
	}

	return popCountByte(num1 ^ num2)
}

func main() {
	sha1 := sha256.Sum256([]byte("x"))
	sha2 := sha256.Sum256([]byte("x"))

	sum := 0
	for index, v1 := range sha1 {
		v2 := sha2[index]
		sum += compare(v1, v2)
	}

	fmt.Println(sum)
}
