package main

import (
	"github.com/cfagiani/aoc2018/util"
	"strings"
	"strconv"
	"fmt"
	"github.com/cfagiani/aoc2018/datastructure"
)

type Operation []int

type State struct {
	before    []int
	operation Operation
	after     []int
}

type Op func(int, int, int, []int)

func main() {
	inputString := util.ReadFileAsString("input/day16.input")
	states, program := processInput(strings.Split(inputString, "\n"))
	funcMappings := part1(states)

	part2(program, mapOpCodes(states, funcMappings))
}

func part1(states []State) [][]Op {
	funcMapping := findMatchingFuncs(states)
	count := 0
	for i := 0; i < len(funcMapping); i++ {
		if len(funcMapping[i]) >= 3 {
			count++
		}
	}
	fmt.Printf("There are %d samples that match 3 or more opcodes\n", count)
	return funcMapping
}

func part2(program []Operation, opCodeMapping map[int]Op) {
	registers := []int{0, 0, 0, 0}
	for i := 0; i < len(program); i++ {
		opCodeMapping[program[i][0]](program[i][1], program[i][2], program[i][3], registers)
	}
	fmt.Printf("After running program, the value in register 0 is: %d", registers[0])
}

func mapOpCodes(states []State, stateMappings [][]Op) map[int]Op {
	codeMappings := make(map[int]*datastructure.Set)

	for i := 0; i < len(stateMappings); i++ {

		curSet := datastructure.NewSet()
		for j := 0; j < len(stateMappings[i]); j++ {
			curSet.Add(stateMappings[i][j])
		}
		if mappingSet, ok := codeMappings[states[i].operation[0]]; ok {
			codeMappings[states[i].operation[0]] = curSet.Union(mappingSet)
		} else {
			codeMappings[states[i].operation[0]] = curSet
		}
	}
	finalMappings := make(map[int]Op)
	for k := range codeMappings {
		for item := range codeMappings[k].Iter() {
			finalMappings[k] = item.(Op)
		}

	}
	return finalMappings
}

func findMatchingFuncs(states []State) [][]Op {
	opcodeFuncs := []Op{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	matches := make([][]Op, len(states))
	for i := 0; i < len(states); i++ {
		matches[i] = testState(states[i], opcodeFuncs)
	}
	return matches
}

func testState(state State, ops []Op) []Op {
	var matches []Op
	for i := 0; i < len(ops); i++ {
		registers := make([]int, 4)
		copy(registers, state.before)
		ops[i](state.operation[1], state.operation[2], state.operation[3], registers)
		if util.IntArrayEquals(registers, state.after) {
			matches = append(matches, ops[i])
		}
	}
	return matches
}

func processInput(lines []string) ([]State, []Operation) {
	var states []State
	var ops []Operation
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		if strings.HasPrefix(lines[i], "Before: ") {
			states = append(states, buildState(lines[i:i+3]))
			i += 3
		} else {
			ops = append(ops, stringToIntArray(lines[i]))
		}
	}
	return states, ops
}

func buildState(lines []string) State {
	return State{before: stringToIntArray(lines[0]),
		operation: stringToIntArray(lines[1]),
		after: stringToIntArray(lines[2])}
}

func stringToIntArray(input string) []int {
	tmp := strings.Replace(input, "Before: [", "", -1)
	tmp = strings.Replace(tmp, "After:  [", "", -1)
	tmp = strings.Replace(tmp, "]", "", -1)
	tmp = strings.Replace(tmp, ",", "", -1)
	parts := strings.Split(tmp, " ")
	results := make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		results[i], _ = strconv.Atoi(parts[i])
	}
	return results
}

func addr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] + registers[b]
}

func addi(a int, b int, c int, registers []int) {
	registers[c] = registers[a] + b
}

func mulr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] * registers[b]
}

func muli(a int, b int, c int, registers []int) {
	registers[c] = registers[a] * b
}

func banr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] & registers[b]
}

func bani(a int, b int, c int, registers []int) {
	registers[c] = registers[a] & b
}

func borr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] | registers[b]
}

func bori(a int, b int, c int, registers []int) {
	registers[c] = registers[a] | b
}

func setr(a int, b int, c int, registers []int) {
	registers[c] = registers[a]
}

func seti(a int, b int, c int, registers []int) {
	registers[c] = a
}

func gtir(a int, b int, c int, registers []int) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a int, b int, c int, registers []int) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a int, b int, c int, registers []int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a int, b int, c int, registers []int) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a int, b int, c int, registers []int) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a int, b int, c int, registers []int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}
