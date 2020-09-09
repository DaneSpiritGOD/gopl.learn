package main

import "fmt"

func distinct(source []string) []string {
	resultTailIndex := 0 //合格序列的尾索引
	for index := 1; index < len(source); index++ {
		if source[index] != source[index-1] {
			resultTailIndex++
			if index != resultTailIndex {
				source[resultTailIndex] = source[index]
			}
		}
	}
	return source[:resultTailIndex+1]
}

func main() {
	list := []string{"a", "b", "b", "c", "c"}
	list = distinct(list)
	fmt.Println(list)

	list = []string{"a", "a", "a", "b", "b"}
	list = distinct(list)
	fmt.Println(list)

	list = []string{"a", "a", "a"}
	list = distinct(list)
	fmt.Println(list)

	list = []string{"a", "b", "b", "b", "c"}
	list = distinct(list)
	fmt.Println(list)
}
