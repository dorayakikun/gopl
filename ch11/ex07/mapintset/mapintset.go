package mapintset

type MapIntSet struct {
	words map[int]bool
}

func (m *MapIntSet) Len() int {
	return len(m.words)
}

func (m *MapIntSet) Add(x int) {
	if m.words == nil {
		m.words = make(map[int]bool)
	}
	m.words[x] = true
}

func (m *MapIntSet) Has(x int) bool {
	return m.words[x]
}

func (m *MapIntSet) Remove(x int) {
	delete(m.words, x)
}

func (m *MapIntSet) Clear() {
	m.words = nil
}

func (m *MapIntSet) UnionWith(t *MapIntSet) {
	if m.words == nil {
		m.words = make(map[int]bool)
	}

	for i := range t.words {
		m.words[i] = true
	}
}