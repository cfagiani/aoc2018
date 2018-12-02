package main

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"github.com/cfagiani/aoc2018/datastructure"
)

func main() {
	inputString := util.ReadFileAsString("input/day1.input")
	lines := strings.Split(inputString, "\n")
	part1(lines)
	var done = false
	var freq = 0
	allFreq := datastructure.NewSet()
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

func part2(lines []string, start int, allFreq *datastructure.Set) (int, bool) {
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
