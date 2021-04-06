package main

import (
	"fmt"

	"github.com/DaneSpiritGOD/ex12.3/sexpr"
)

type movie struct {
	Title, Subtitle string
	Year            int
	Ids             []int
	Color1          bool
	Color2          bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	Distance1       float32
	Distance2       complex64
	Foo             interface{}
	Bar             interface{}
}

func main() {
	strangelove := movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Ids:      []int{1, 3, 5, 7, 9},
		Color1:   false,
		Color2:   true,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Distance1: 3.1415926,
		Distance2: 3 + 1i,
		Foo:       [3]int{8, 5, 6},
		Bar:       map[int]string{1: "a", 2: "b"},
	}

	content, err := sexpr.Marshal(strangelove)
	if err != nil {
		fmt.Printf("Marshal error: %v", err)
	}

	fmt.Println(string(content))
}
