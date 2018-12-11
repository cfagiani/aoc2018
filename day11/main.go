package main

import (
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"fmt"
)

func main() {
	inputString := util.ReadFileAsString("input/day11.input")
	serialNum, _ := strconv.Atoi(inputString)
	grid := populatePower(serialNum)
	levelGrid := populateCells(grid)
	part1(levelGrid)
}


func part1(cells [296][296]int) {
	maxX := 0
	maxY := 0
	for i := 0; i < len(cells); i++ {
		for j := 0; j < len(cells[i]); j++ {
			if cells[i][j] > cells[maxX][maxY] {
				maxX = i
				maxY = j
			}
		}
	}
	fmt.Printf("Max grid starts at (%d,%d)\n", maxX+1, maxY+1)
}
func populateCells(grid [300][300]int) [296][296]int {
	powerCells := [296][296] int{}
	for i := 0; i < len(powerCells); i++ {
		for j := 0; j < len(powerCells[i]); j++ {
			powerCells[i][j] = getGridPower(grid, i, j, 3)
		}
	}
	return powerCells
}

func getGridPower(grid [300][300]int, x int, y int, size int) int {
	sum := 0
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			sum += grid[i][j]
		}
	}
	return sum
}

func populatePower(serialNum int) [300][300]int {
	grid := [300][300] int{}
	for i := 0; i < 300; i++ {
		for j := 0; j < 300; j++ {
			rackId := getRackId(i, j)
			power := rackId * (1 + j)
			power += serialNum
			power *= rackId
			grid[i][j] = getDigit(power) - 5
		}
	}
	return grid
}

func getRackId(x int, y int) int {
	return x + 1 + 10
}

func getDigit(val int) int {
	return (val / 100) % 10
}
