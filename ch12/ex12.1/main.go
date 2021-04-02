package main

import "github.com/DaneSpiritGOD/ex12.1/display"

type extraKey1 struct {
	Number int
	Author string
}

type movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string

	Extra1 map[extraKey1]int
	Extra2 map[[3]string]int
}

type cycle struct {
	Value int
	Tail  *cycle
}

func main() {
	strangelove := movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
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

		Extra1: map[extraKey1]int{
			{1, "Dane"}: 9,
			{2, "Fei"}:  9,
		},
		Extra2: map[[3]string]int{
			{"1", "2", "3"}: 10,
			{"4", "5", "6"}: 11,
		},
	}

	display.Display("movie", strangelove)

	var c cycle
	c = cycle{42, &c}
	display.Display("c", c)
}
