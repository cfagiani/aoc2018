package datastructure

var exists = struct{}{}

type Set struct {
	set map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{make(map[interface{}]struct{})}
}

func NewSetWith(items ...interface{}) *Set {
	s := NewSet()
	for i := 0; i < len(items); i++ {
		s.Add(items[i])
	}
	return s
}

func (set *Set) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = exists
	return !found //False if it existed already
}

func (set *Set) Contains(i interface{}) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *Set) Remove(i interface{}) {
	delete(set.set, i)
}

func (set *Set) Clear() {
	*set = *NewSet()
}

func (set *Set) Size() int {
	return len(set.set)
}

func (set *Set) Union(other *Set) *Set {
	newSet := NewSet()
	for item, _ := range set.set {
		newSet.Add(item)
	}
	for item, _ := range other.set {
		newSet.Add(item)
	}
	return newSet
}

func (set *Set) Intersect(other *Set) Set {
	newSet := NewSet()
	for item, _ := range set.set {
		if other.Contains(item) {
			newSet.Add(item)
		}
	}
	return *newSet
}

func (set *Set) Difference(other *Set) Set {
	newSet := NewSet()
	for item, _ := range set.set {
		if !other.Contains(item) {
			newSet.Add(item)
		}
	}
	return *newSet
}

//Apply test function to each element in the set. If the test returns true, stop iteration.
func (set *Set) Each(test func(interface{}) bool) {
	for elem, _ := range set.set {
		if test(elem) {
			break
		}
	}
}

//Obtain an iterator for the set so we can use it in range
func (set *Set) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		for elem, _ := range set.set {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}
