package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"strings"
)

func main() {
	inputString := util.ReadFileAsString("input/day17.input")
	state := processInput(strings.Split(inputString, "\n"))
	part1(state)
	part2(state)
}

func part1(state [][]byte) {
	fill(state, 500, 0)
	fmt.Printf("Water can reach %d locs\n", countWet(state, false))
}

func part2(state [][]byte) {
	fmt.Printf("Water retained: %d \n", countWet(state, true))
}

func fill(state [][]byte, x int, y int) {
	if y >= len(state)-1 || !isOpen(x, y, state) {
		return
	}

	if !isOpen(x, y+1, state) {
		leftX := x
		for isOpen(leftX, y, state) && !isOpen(leftX, y+1, state) {
			state[y][leftX] = '|'
			leftX--
		}
		rightX := x + 1
		for isOpen(rightX, y, state) && !isOpen(rightX, y+1, state) {
			state[y][rightX] = '|'
			rightX++
		}
		if isOpen(leftX, y+1, state) || isOpen(rightX, y+1, state) {
			fill(state, leftX, y)
			fill(state, rightX, y)
		} else if state[y][leftX] == '#' && state[y][rightX] == '#' {
			for x2 := leftX + 1; x2 < rightX; x2++ {
				state[y][x2] = '~'
			}
		}
	} else if state[y][x] == 0 {
		state[y][x] = '|'
		fill(state, x, y+1)
		if state[y+1][x] == '~' {
			fill(state, x, y)
		}
	}
}

func countWet(state [][]byte, waterOnly bool) int {
	wetCount := 0
	waterCount := 0
	for i := getMinY(state); i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			if !waterOnly && state[i][j] == '|' {
				wetCount++
			} else if state[i][j] == '~' {
				waterCount++
			}
		}
	}
	return waterCount + wetCount
}

func getMinY(state [][]byte) int {
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			if state[i][j] == '#' {
				return i
			}
		}
	}
	return -1
}

func processInput(inputLines []string) [][]byte {
	maxX, maxY := 0, 0
	for i := 0; i < len(inputLines); i++ {
		_, toX, _, toY := getCoordinates(inputLines[i])
		if toX > maxX {
			maxX = toX
		}
		if toY > maxY {
			maxY = toY
		}
	}
	state := make([][]byte, maxY+2)
	for i := 0; i < maxY+2; i++ {
		state[i] = make([]byte, maxX+2)
	}
	for i := 0; i < len(inputLines); i++ {
		fromX, toX, fromY, toY := getCoordinates(inputLines[i])
		for y := fromY; y <= toY; y++ {
			for x := fromX; x <= toX; x++ {
				state[y][x] = '#'
			}
		}
	}
	return state
}

func getCoordinates(line string) (int, int, int, int) {
	parts := strings.Split(line, ", ")
	if strings.HasPrefix(parts[0], "x=") {
		x1, x2 := getNumRange(parts[0])
		y1, y2 := getNumRange(parts[1])
		return x1, x2, y1, y2
	} else {
		x1, x2 := getNumRange(parts[1])
		y1, y2 := getNumRange(parts[0])
		return x1, x2, y1, y2
	}
}

func getNumRange(line string) (int, int) {
	start, end := 0, 0
	numStr := strings.Split(line, "=")[1]
	parts := strings.Split(numStr, "..")
	if len(parts) > 1 {
		start, _ = strconv.Atoi(parts[0])
		end, _ = strconv.Atoi(parts[1])
	} else {
		start, _ = strconv.Atoi(parts[0])
		end = start
	}
	return start, end
}

func isOpen(x int, y int, state [][]byte) bool {
	return state[y][x] == 0 || state[y][x] == '|'
}
