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
		fmt.Printf("%s\n", state.toString(true))
	}
	sum := state.getSumOfIndexes()
	fmt.Printf("Sum of all plants: %d\n", sum)
}

func applyGeneration(state *State, rules []Rule) *State {
	var nextPots []bool
	//starts nextPots off with 2 false entries since we won't look at those in the loop
	for i := 0; i < NeighborCount; i++ {
		nextPots = append(nextPots, false)
	}

	//current state's pots also needs to be padded on both ends so we don't get an outOfBounds error
	curPots := append([]bool{false, false}, state.pots...)
	curPots = append(curPots, []bool{false, false}...)

	for i := NeighborCount; i < len(curPots)-NeighborCount; i++ {
		targetPattern := getTargetPattern(curPots, i)
		matched := false
		for j := 0; j < len(rules); j++ {
			if rules[j].matches(targetPattern) {
				nextPots = append(nextPots, rules[j].result)
				matched = true
				break
			}
		}
		if !matched {
			nextPots = append(nextPots, false)
		}
	}
	for i := 0; i < NeighborCount; i++ {
		nextPots = append(nextPots, false)
	}
	//now we can trim any negative positions up to the first plant - the 2 positions we added
	trimSize := 0
	for trimSize = 0; trimSize < state.zeroIndex-NeighborCount; trimSize++ {
		if nextPots[trimSize] {
			break
		}
	}
	nextPots = nextPots[trimSize:]
	//can also trim off the other end
	lastIdx := 0
	for lastIdx = len(nextPots) - 1; lastIdx > 0; lastIdx-- {
		if nextPots[lastIdx] {
			break
		}
	}
	nextPots = nextPots[:lastIdx+NeighborCount]
	//since we pre-pended NeighborCount entries, advance the zero index by that much
	nextState := State{zeroIndex: state.zeroIndex + NeighborCount - trimSize, pots: nextPots}
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
	for i := startIdx - NeighborCount; i <= startIdx+NeighborCount; i++ {
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
