package charCount_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/DaneSpiritGOD/ex11.1/charCount"
)

func TestCount(t *testing.T) {
	var tests = []struct {
		input   string
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{"abc", map[rune]int{'a': 1, 'b': 1, 'c': 1}, []int{0, 3, 0, 0, 0}, 0},
		{"aabbcc", map[rune]int{'a': 2, 'b': 2, 'c': 2}, []int{0, 6, 0, 0, 0}, 0},
	}

	for _, test := range tests {
		counts, utflen, invalid := charCount.Count(strings.NewReader(test.input))

		if !reflect.DeepEqual(test.counts, counts) {
			t.Errorf("charCount(%s) counts = %v, expected: %v", test.input, counts, test.counts)
		}

		if !reflect.DeepEqual(test.utflen, utflen) {
			t.Errorf("charCount(%s) utflen = %v, expected: %v", test.input, utflen, test.utflen)
		}

		if invalid != test.invalid {
			t.Errorf("charCount(%s) invalid = %v, expected: %v", test.input, invalid, test.invalid)
		}
	}
}
