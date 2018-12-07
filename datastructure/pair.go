package datastructure

import "sort"

type IntPair struct {
	Key   int
	Value int
}

type IntPairList []IntPair

func (p IntPairList) Len() int           { return len(p) }
func (p IntPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p IntPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//Sorts a map with int values and returns as a list of pairs
func SortMapByValue(input map[int]int, reverse bool) IntPairList {
	pl := make(IntPairList, len(input))
	i := 0
	for k, v := range input {
		pl[i] = IntPair{k, v}
		i++
	}
	if reverse {
		sort.Sort(sort.Reverse(pl))
	} else {
		sort.Sort(pl)
	}
	return pl
}
