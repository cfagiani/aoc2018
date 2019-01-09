package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strings"
)

func main() {
	inputString := util.ReadFileAsString("input/day18.input")
	state := processInput(strings.Split(inputString, "\n"))
	part1(state)
	part2(processInput(strings.Split(inputString, "\n")))
}

func part1(state [][]byte) {
	for i := 0; i < 10; i++ {
		state = getNextState(state)
	}
	lumberYards := count(state, '#')
	woods := count(state, '|')
	fmt.Printf("Resource value after 10 minutes is %d\n", lumberYards*woods)
}

func part2(state [][]byte) {
	var prevVals []int
	target := 1000000000

	for i := 0; i < target; i++ {
		state = getNextState(state)
		lumberYards := count(state, '#')
		woods := count(state, '|')
		val := lumberYards * woods
		repeatValIdx := -1

		for j := 0; j < len(prevVals); j++ {
			if val == prevVals[j] {
				repeatValIdx = j
				if repeatValIdx < len(prevVals) {
					peekState := getNextState(state)
					if count(peekState, '#')*count(peekState, '|') != prevVals[repeatValIdx+1] {
						//not really a pattern
						repeatValIdx = -1
					} else {
						break
					}
				}
			}
		}
		if repeatValIdx != -1 {
			pattern := prevVals[repeatValIdx:]

			remainingIters := target - i - 1
			fmt.Printf("Resource value after %d minutes is %d\n", target, pattern[remainingIters%len(pattern)])
			return
		}
		prevVals = append(prevVals, val)

	}

}

func getNextState(state [][]byte) [][]byte {
	nextState := make([][]byte, len(state))
	for i := 0; i < len(state); i++ {
		nextState[i] = make([]byte, len(state[i]))
		for j := 0; j < len(state[i]); j++ {
			nextState[i][j] = getNextStateForAcre(i, j, state)
		}
	}
	return nextState
}

func getNextStateForAcre(i int, j int, state [][]byte) byte {
	switch state[i][j] {
	case '.':
		if countAdjacent(i, j, state, '|') >= 3 {
			return '|'
		}
	case '|':
		if countAdjacent(i, j, state, '#') >= 3 {
			return '#'
		}
	case '#':
		if countAdjacent(i, j, state, '#') >= 1 && countAdjacent(i, j, state, '|') >= 1 {
			return '#'
		} else {
			return '.'
		}
	}
	return state[i][j]
}

func count(state [][]byte, acreType byte) int {
	count := 0
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			if state[i][j] == acreType {
				count++
			}
		}
	}
	return count
}

func countAdjacent(i int, j int, state [][]byte, acreType byte) int {
	count := 0
	for a := -1; a < 2; a++ {
		for b := -1; b < 2; b++ {
			if i+a >= 0 && i+a < len(state) {
				if j+b >= 0 && j+b < len(state[i+a]) {
					if !(a == 0 && b == 0) && state[i+a][j+b] == acreType {
						count++
					}
				}
			}
		}
	}
	return count
}

func processInput(lines []string) [][]byte {
	state := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		state[i] = make([]byte, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			state[i][j] = lines[i][j]
		}
	}
	return state
}
