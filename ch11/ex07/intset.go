package intset

import (
	"bytes"
	"fmt"
	"math/bits"
)

type IntSet struct {
	words []uint
}

// uintが64bitで計算された時、(^uint(0) >> 63)の結果は1となる
// そのため、32を1bit左シフトすることで最終的な解は64
// uintが32bitで計算された時、左シフトが発生しないので32のまま
const d = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/d, uint(x%d)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/d, uint(x%d)
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
		for j := 0; j < d; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", d*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var bitcount int
	for _, word := range s.words {
		bitcount += bits.OnesCount(word)
	}
	return bitcount
}

func (s *IntSet) Remove(x int) {
	word, bit := x/d, uint(x%d)
	s.words[word] &= ^uint(1 << bit)
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	return &IntSet{words: append([]uint(nil), s.words...)}
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	a := len(s.words)
	b := len(t.words)

	if a == 0 || b == 0 {
		return
	}

	if b < 2 {
		s.words[0] &= t.words[0]
		return
	}

	var c int
	if a < b {
		c = a
	} else {
		c = b
	}

	for i := 0; i < c; i++ {
		s.words[i] &= t.words[i]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	a := len(s.words)
	b := len(t.words)

	if a == 0 || b == 0 {
		return
	}

	if b < 2 {
		s.words[0] = (s.words[0] | t.words[0]) ^ t.words[0]
		return
	}

	var c int
	if a < b {
		c = a
	} else {
		c = b
	}

	for i := 0; i < c; i++ {
		s.words[i] = (s.words[i] | t.words[i]) ^ t.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	a := len(s.words)
	b := len(t.words)

	if a == 0 || b == 0 {
		return
	}

	if b < 2 {
		s.words[0] ^= t.words[0]
		return
	}

	// 左辺のwordsが小さい場合は、右辺と長さを合わせる
	// ※右辺側の値を漏れなく反映したいので
	if a < b {
		delta := b - a
		for i := 0; i < delta; i++ {
			s.words = append(s.words, 0)
		}
	}

	for i := 0; i < b; i++ {
		s.words[i] ^= t.words[i]
	}
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < d; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, d*i+j)
			}
		}
	}
	return elems
}
