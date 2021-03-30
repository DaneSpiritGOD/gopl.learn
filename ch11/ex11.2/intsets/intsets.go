package intsets

import (
	"bytes"
	"fmt"
)

const bitLen = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitLen, uint(x%bitLen)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/bitLen, uint(x%bitLen)
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

// IntersectWith 交集：元素在A集合B集合均出现
func (s *IntSet) IntersectWith(t *IntSet) IntSet {
	ss := IntSet{}

	for i, tword := range t.words {
		if i >= len(s.words) {
			break
		}

		ss.words = append(ss.words, s.words[i]&tword)
	}

	return ss
}

// DifferenceWith 差集：元素出现在A集合，未出现在B集合
func (s *IntSet) DifferenceWith(t *IntSet) IntSet {
	ss := IntSet{}

	for i, tword := range s.words {
		if i >= len(t.words) {
			ss.words = append(ss.words, tword)
		} else {
			ss.words = append(ss.words, tword&^t.words[i])
		}
	}

	return ss
}

// SymmetricDifference 并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A
func (s *IntSet) SymmetricDifference(t *IntSet) IntSet {
	ss := IntSet{}

	i := 0
	for ; i < len(s.words); i++ {
		sword := s.words[i]

		if i >= len(t.words) {
			ss.words = append(ss.words, sword)
		} else {
			tword := t.words[i]
			ss.words = append(ss.words, sword^tword)
		}
	}

	for ; i < len(t.words); i++ {
		tword := t.words[i]
		ss.words = append(ss.words, tword)
	}

	return ss
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitLen; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitLen*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popCount(x uint) int {
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
	word, bit := x/bitLen, uint(x%bitLen)
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
	ss.words = append(ss.words, s.words...)
	return &ss
}

// Elems return all elements of IntSet
func (s *IntSet) Elems() []int {
	var buf []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitLen; j++ {
			if word&(1<<uint(j)) != 0 {
				buf = append(buf, bitLen*i+j)
			}
		}
	}
	return buf
}
