package main

import (
	"github.com/cfagiani/aoc2018/util"
	"strings"
	"fmt"
)

type Rule struct {
	result  bool
	pattern []bool
}

type State struct {
	zeroIndex int
	pots      []bool
}

const NeighborCount = 2

func main() {
	inputString := util.ReadFileAsString("input/day12.input")
	initState, rules := initialize(inputString)
	fmt.Printf("%s\n", initState.toString(false))
	part1(&initState, rules) //2054 too high
}

func part1(state *State, rules []Rule) {
	for i := 0; i < 20; i++ {
		state = applyGeneration(state, rules)
		fmt.Printf("%s\n", state.toString(false))
	}
	sum := state.getSumOfIndexes()

	fmt.Printf("Sum of all plants: %d\n", sum)
}

func applyGeneration(state *State, rules []Rule) *State {
	var pots []bool
	for i := 0; i < NeighborCount; i++ {
		pots = append(pots, false)
	}
	var nextState State
	nextState.zeroIndex = state.zeroIndex + NeighborCount
	curPots := append([]bool{false, false}, state.pots...)
	curPots = append(curPots, []bool{false, false}...)

	for i := 2; i < len(curPots)-NeighborCount; i++ {
		targetPattern := getTargetPattern(curPots, i)
		matched := false
		for j := 0; j < len(rules); j++ {
			if rules[j].matches(targetPattern) {
				pots = append(pots, rules[j].result)
				matched = true
				break
			}
		}
		if !matched {
			pots = append(pots, false)
		}
	}
	for i := 0; i < NeighborCount; i++ {
		pots = append(pots, false)
	}
	nextState.pots = pots
	return &nextState
}

func (state State) toString(withCaret bool) string {
	var str string
	for i := 0; i < len(state.pots); i++ {
		if state.pots[i] {
			str += "#"
		} else {
			str += "."
		}
	}
	if withCaret {
		str += "\n"
		for i := 0; i < state.zeroIndex; i++ {
			str += " "
		}
		str += "^"
	}
	return str
}

func (state State) getSumOfIndexes() int {
	sum := 0
	for i := 0; i < len(state.pots); i++ {
		if state.pots[i] {
			sum += i - state.zeroIndex
		}
	}
	return sum
}

func (rule Rule) matches(pattern []bool) bool {
	for i := 0; i < len(pattern); i++ {
		if rule.pattern[i] != pattern[i] {
			return false
		}
	}
	return true
}

func getTargetPattern(state []bool, startIdx int) []bool {
	var pattern []bool
	for i := startIdx - NeighborCount; i < startIdx+NeighborCount+1; i++ {
		pattern = append(pattern, state[i])

	}

	return pattern
}

func initialize(input string) (State, []Rule) {
	lines := strings.Split(input, "\n")
	initialState := getInitState(lines[0])
	rules := initRules(lines[2:])
	return initialState, rules
}

func initRules(lines []string) []Rule {
	var rules []Rule
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " => ")
		rules = append(rules, Rule{result: buildBoolArray(parts[1])[0], pattern: buildBoolArray(parts[0])})
	}
	return rules
}

func getInitState(line string) State {
	parts := strings.Split(line, ": ")
	return State{pots: buildBoolArray(parts[1]), zeroIndex: 0}
}

func buildBoolArray(line string) []bool {
	var state []bool
	for i := 0; i < len(line); i++ {
		if line[i] == '#' {
			state = append(state, true)
		} else {
			state = append(state, false)
		}
	}
	return state
}
