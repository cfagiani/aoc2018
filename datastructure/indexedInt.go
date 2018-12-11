package datastructure

import "sort"

type IndexedInt struct {
	Key   int
	Value int
}

type IndexedIntList []IndexedInt

func (p IndexedIntList) Len() int           { return len(p) }
func (p IndexedIntList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p IndexedIntList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//Sorts a map with int values and returns as a list of pairs
func SortMapByValue(input map[int]int, reverse bool) IndexedIntList {
	pl := make(IndexedIntList, len(input))
	i := 0
	for k, v := range input {
		pl[i] = IndexedInt{k, v}
		i++
	}
	if reverse {
		sort.Sort(sort.Reverse(pl))
	} else {
		sort.Sort(pl)
	}
	return pl
}
