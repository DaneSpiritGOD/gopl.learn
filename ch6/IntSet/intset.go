package intset

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds the non-negative values items to the set.
func (s *IntSet) AddAll(items ...int) {
	for _, item := range items {
		s.Add(item)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popCount(x uint64) int {
	var sum int = 0
	for x != 0 {
		x = x & (x - 1)
		sum++
	}
	return sum
}

// Equals return two IntSet are euqual to each other
func (s *IntSet) Equals(s1 *IntSet) bool {
	if len(s.words) != len(s1.words) {
		return false
	}

	for i, word := range s.words {
		word1 := s1.words[i]
		if word != word1 {
			return false
		}
	}

	return true
}

// Len return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popCount(word)
	}
	return count
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) && s.words[word]&(1<<bit) != 0 {
		s.words[word] &^= 1 << bit
	}
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	ss := IntSet{}
	for _, word := range s.words {
		ss.words = append(ss.words, word)
	}

	return &ss
}
