package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

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

func (s *IntSet) Len() int {
	//var count int
	//for i, word := range s.words {
	//	if word == 0 {
	//		continue
	//	}
	//	for j := 0; j < 64; j++ {
	//		if word&(1<<uint(j)) != 0 {
	//			count += 64*i+j
	//		}
	//	}
	//}
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	fmt.Println(s.words)
	word, bit := x/64, uint(x%64)
	s.words[word] &= ^uint64(1 << bit)
	if s.words[word] == 0 {

	}
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{ words: append([]uint64(nil), s.words...)}
}

