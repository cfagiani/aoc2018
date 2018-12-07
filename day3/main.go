package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"strings"
)

func main() {
	inputString := util.ReadFileAsString("input/day3.input")
	lines := strings.Split(inputString, "\n")
	overlapCount := buildOverlapCount(lines)
	inchCount := part1(overlapCount)
	fmt.Printf("%d inches overlap\n", inchCount)
	nonOverlap := part2(overlapCount, lines)
	fmt.Printf("Square %s has no overlap\n", nonOverlap)
}

//counts the number of square inches occupied by more than 1 square.
func part1(counts [][]int) int {
	inchCount := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if counts[i][j] > 1 {
				inchCount++
			}
		}
	}
	return inchCount
}

//Finds the ID of the square with no overlaps
func part2(counts [][]int, lines []string) string {
	for i := 0; i < len(lines); i++ {
		id, startX, startY, w, h := tokenizeLine(lines[i])
		if checkSquare(counts, startX, startY, w, h) {
			return id
		}
	}
	return ""
}

//Checks if a square is the only one that occupies a space. If so, return true, otherwise false
func checkSquare(counts [][]int, startX int, startY int, w int, h int) bool {
	for x := startX; x < startX+w; x++ {
		for y := startY; y < startY+h; y++ {
			if counts[x][y] > 1 {
				return false
			}
		}
	}
	return true
}

//Builds count of how many squares occupy a square inch
func buildOverlapCount(lines []string) [][]int {
	overlap := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		overlap[i] = make([]int, 1000)
	}

	for i := 0; i < len(lines); i++ {
		_, startX, startY, w, h := tokenizeLine(lines[i])
		for x := startX; x < startX+w; x++ {
			for y := startY; y < startY+h; y++ {
				overlap[x][y]++
			}
		}
	}
	return overlap
}

//breaks a line of the format #1185 @ 145,202: 22x13 into the following fields:
//id, startX, startY, width, height
func tokenizeLine(line string) (string, int, int, int, int) {
	parts := strings.Split(line, "@")
	dimensions := strings.Split(parts[1], ":")
	startPos := strings.Split(strings.Trim(dimensions[0], " "), ",")

	hw := strings.Split(strings.Trim(dimensions[1], " "), "x")
	startX, _ := strconv.Atoi(startPos[0])
	startY, _ := strconv.Atoi(startPos[1])
	w, _ := strconv.Atoi(hw[0])
	h, _ := strconv.Atoi(hw[1])
	return strings.Trim(parts[0], " "), startX, startY, w, h
}
