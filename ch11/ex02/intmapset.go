package intset

type IntMapSet struct {
	v map[int]bool
}

func (i *IntMapSet) Add(x int) {
	i.v[x] = true
}

func (i *IntMapSet) Has(x int) bool {
	return i.v[x]
}
