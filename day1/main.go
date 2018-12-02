package main

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/cfagiani/aoc2018/util"
)

func main() {
	inputString := util.ReadFileAsString("input/day1.input")
	lines := strings.Split(inputString, "\n")
	part1(lines)
	var done = false
	var freq = 0
	allFreq := NewIntSet()
	for (!done) {
		freq, done = part2(lines, freq, allFreq)
	}
	fmt.Printf("First dupe is %d\n", freq)
}

func part1(lines []string) {
	sum := 0
	for _, val := range lines {
		intVal, _ := strconv.Atoi(val)
		sum += intVal
	}

	fmt.Printf("Final frequency is %d\n", sum)
}

func part2(lines []string, start int, allFreq *IntSet) (int, bool) {
	sum := start
	for _, val := range lines {
		intVal, _ := strconv.Atoi(val)
		sum += intVal
		if ! allFreq.Add(sum) {
			return sum, true
		}
	}
	return sum, false
}

type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *IntSet) Contains(i int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *IntSet) Remove(i int) {
	delete(set.set, i)
}

func (set *IntSet) Size() int {
	return len(set.set)
}
