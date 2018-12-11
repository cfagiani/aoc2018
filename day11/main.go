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
	part1(grid)
	part2(grid)
}

func part2(cells [][]int) {
	maxVal := 0
	maxX := 0
	maxY := 0
	maxSize := 0
	for i := 1; i < len(cells); i++ {
		x, y, val := getMaxGrid(cells, i)
		if val > maxVal {
			maxVal = val
			maxX = x
			maxY = y
			maxSize = i
		}
	}
	fmt.Printf("Max grid id is %d,%d,%d\n", maxX, maxY, maxSize)
}

func part1(grid [][]int) {
	maxX, maxY, _ := getMaxGrid(grid, 3)
	fmt.Printf("Max grid starts at (%d,%d)\n", maxX, maxY)
}

func getMaxGrid(grid [][]int, size int) (int, int, int) {
	cells := populateCells(grid, size)
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
	return maxX + 1, maxY + 1, cells[maxX][maxY]
}

func populateCells(grid [][]int, size int) [][]int {
	powerCells := initGrid(len(grid) - size + 1)
	for i := 0; i < len(powerCells); i++ {
		for j := 0; j < len(powerCells[i]); j++ {
			powerCells[i][j] = getGridPower(grid, i, j, size)
		}
	}
	return powerCells
}

func getGridPower(grid [][]int, x int, y int, size int) int {
	sum := 0
	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			sum += grid[i][j]
		}
	}
	return sum
}

func populatePower(serialNum int) [][]int {
	grid := initGrid(300)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
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

func initGrid(size int) [][]int {
	arr := make([][]int, size)
	for i := range arr {
		arr[i] = make([]int, size)
	}
	return arr
}
