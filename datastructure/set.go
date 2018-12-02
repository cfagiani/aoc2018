package datastructure

var exists = struct{}{}

type Set struct {
	set map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{make(map[interface{}]struct{})}
}

func (set *Set) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = exists
	return !found //False if it existed already
}

func (set *Set) Contains(i int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *Set) Remove(i interface{}) {
	delete(set.set, i)
}

func (set *Set) Size() int {
	return len(set.set)
}
