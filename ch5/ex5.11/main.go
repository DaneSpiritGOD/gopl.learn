package main

import (
	"fmt"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, o := range topSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, o)
	}
}

// MAP _
type MAP map[string][]string

func topSort(m MAP) []string {
	var order []string
	seen := make(map[string]bool)

	var visit2 func(items []string)
	visit2 = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visit2(m[item])
				order = append(order, item)
			} else {
				// seen[item]==true, but not in order
				for o := range order {
					if order[o] == item {
						panic(fmt.Sprintf("circle detected: %s", item))
					}
				}
			}
		}
	}

	for k := range m {
		items := [1]string{k}
		visit2(items[:])
	}
	return order
}
