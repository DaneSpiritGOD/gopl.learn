package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track type of music track
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type key2Sort struct {
	t []*Track
	// less func(x, y *Track) bool
	key1 string
	key2 string
}

func (x key2Sort) Len() int { return len(x.t) }
func (x key2Sort) Less(i, j int) bool {
	xi := x.t[i]
	xj := x.t[j]

	swi := func(k string) (result bool, next bool) {
		switch k {
		case "Title":
			if xi.Title != xj.Title {
				return xi.Title < xj.Title, false
			}
		case "Artist":
			if xi.Artist != xj.Artist {
				return xi.Artist < xj.Artist, false
			}
		case "Album":
			if xi.Album != xj.Album {
				return xi.Album < xj.Album, false
			}
		case "Year":
			if xi.Year != xj.Year {
				return xi.Year < xj.Year, false
			}
		case "Length":
			if xi.Length != xj.Length {
				return xi.Length < xj.Length, false
			}
		default:
			panic("error of key")
		}
		return false, true
	}

	kr, next := swi(x.key1)
	if !next {
		return kr
	}

	kr, next = swi(x.key2)
	if !next {
		return kr
	}

	return false
}

func (x key2Sort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	// sort.Sort(byArtist(tracks))
	// printTracks(tracks)

	sort.Sort(key2Sort{tracks, "Year", "Title"})
	printTracks(tracks)
}
