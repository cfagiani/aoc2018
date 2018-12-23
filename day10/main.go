package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/datastructure"
	"github.com/cfagiani/aoc2018/util"
	"math"
	"strconv"
	"strings"
)

type Star struct {
	pos      *datastructure.IntPair
	velocity *datastructure.IntPair
}

const PositionLabel = "position=<"
const VelocityLabel = "velocity=<"
const MaxSeconds = 30000

func main() {
	inputString := util.ReadFileAsString("input/day10.input")
	stars := buildStars(inputString)
	seconds := part1(stars)
	part2(seconds)
}

func part2(val int) {
	fmt.Printf("Took %d seconds", val+1)
}

func part1(stars []Star) int {
	minVal := 99999
	minIdx := -1
	for i := 0; i < MaxSeconds; i++ {
		for j := 0; j < len(stars); j++ {
			stars[j].pos = calcNewPos(stars[j])

		}
		min, max := computeBounds(stars)
		if max.B-min.B < 20 {
			if max.B-min.B < minVal {
				minIdx = i
				minVal = max.B - min.B
			}
			fmt.Printf("%d:\n", i)
			normalizeAndPrint(stars)
		}
	}
	return minIdx
}

func normalizeAndPrint(stars []Star) {

	min, max := computeBounds(stars)
	gridSizeX := 60
	gridSizeY := 12
	var grid [12][60]bool
	for i := 0; i < len(stars); i++ {
		normX, normY := normalizePair(*stars[i].pos, min, max)
		x, y := convertToIndexes(normX, normY, float64(gridSizeX), float64(gridSizeY))
		grid[y][x] = true
	}

	///now print
	for i := 0; i < gridSizeY; i++ {
		for j := 0; j < gridSizeX; j++ {
			if grid[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n\n\n")
}

func convertToIndexes(x float64, y float64, gridSizeX float64, gridSizeY float64) (int, int) {
	return int(math.Max(0, x*gridSizeX-1)), int(math.Max(0, y*gridSizeY-1))
}

func computeBounds(stars []Star) (datastructure.IntPair, datastructure.IntPair) {
	minX := 100000
	minY := 100000
	maxX := 0
	maxY := 0
	for i := 0; i < len(stars); i++ {
		if stars[i].pos.A > maxX {
			maxX = stars[i].pos.A
		}
		if stars[i].pos.B > maxY {
			maxY = stars[i].pos.B
		}
		if stars[i].pos.A < minX {
			minX = stars[i].pos.A
		}
		if stars[i].pos.B < minY {
			minY = stars[i].pos.B
		}
	}
	return datastructure.IntPair{A: minX, B: minY}, datastructure.IntPair{A: maxX, B: maxY}
}

func normalizePair(pair datastructure.IntPair, min datastructure.IntPair, max datastructure.IntPair) (float64, float64) {
	return float64(pair.A-min.A) / float64(max.A-min.A), float64(pair.B-min.B) / float64(max.B-min.B)
}

func buildStars(input string) []Star {
	var stars []Star
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		stars = append(stars, buildStar(lines[i]))
	}
	return stars
}

func buildStar(input string) Star {
	positionStr := input[len(PositionLabel) : strings.Index(input, VelocityLabel)-2]
	velStr := input[strings.Index(input, VelocityLabel)+len(VelocityLabel) : len(input)-1]

	return Star{pos: buildPair(positionStr), velocity: buildPair(velStr)}
}

func buildPair(input string) *datastructure.IntPair {
	parts := strings.Split(input, ", ")
	a, _ := strconv.Atoi(strings.Replace(parts[0], " ", "", -1))
	b, _ := strconv.Atoi(strings.Replace(parts[1], " ", "", -1))
	return &datastructure.IntPair{A: a, B: b}
}

func calcNewPos(s Star) *datastructure.IntPair {
	return &datastructure.IntPair{A: s.pos.A + s.velocity.A, B: s.pos.B + s.velocity.B}
}
