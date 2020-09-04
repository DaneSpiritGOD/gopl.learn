package main

import "fmt"

func rotate(data []int, count int) {
	if count == 0 {
		return
	}

	if count > 0 {
		var temp []int
		temp = append(temp, data[:count]...)
		copy(data[:], data[count:])
		copy(data[len(data)-count:], temp)
		return
	}

	rotate(data, len(data)+count)
}

func main() {
	data := [...]int{1, 2, 3, 4, 5, 6}
	rotate(data[:], 2)
	fmt.Println(data)

	data = [...]int{1, 2, 3, 4, 5, 6}
	rotate(data[:], -2)
	fmt.Println(data)
}
