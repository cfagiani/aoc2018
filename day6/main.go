package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/datastructure"
	"github.com/cfagiani/aoc2018/util"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	inputString := util.ReadFileAsString("input/day6.input")
	lines := strings.Split(inputString, "\n")
	points := buildPointList(lines)
	part1(points)
	part2(points)
}

func part2(points []Point) {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			sum := 0
			for _, v := range points {
				sum += manhattanDist(i, j, v.x, v.y)
			}
			if sum < 10000 {
				count++
			}
		}
	}
	fmt.Printf("size of area %d", count)
}

func part1(points []Point) {
	minX, minY, maxX, maxY := findBounds(points)
	area := make([]int, len(points))
	boundsIndexes := getIndexesOfBounds(points, minX, minY, maxX, maxY)
	for i := minX; i < maxX; i++ {
		for j := minY; j < maxY; j++ {
			closestIdx := findIndexOfClosestPoint(points, i, j)
			if closestIdx >= 0 {
				if !util.IsIntInSlice(closestIdx, boundsIndexes) {
					area[closestIdx]++
				}
			}
		}
	}

	//this is a hacky way to remove the "infinite" regions. Should really do something better
	area2 := make([]int, len(points))
	for i := minX; i < maxX*2; i++ {
		for j := minY; j < maxY*2; j++ {
			closestIdx := findIndexOfClosestPoint(points, i, j)
			if closestIdx >= 0 {
				if !util.IsIntInSlice(closestIdx, boundsIndexes) {
					area2[closestIdx]++
				}
			}
		}
	}

	for i := 0; i < len(area); i++ {
		if area[i] != area2[i] {
			area[i] = 0
		}
	}

	maxArea := 0
	for i := 0; i < len(area); i++ {
		if area[i] > maxArea {
			maxArea = area[i]
		}
	}
	fmt.Printf("Max area is %d\n", maxArea)
}

func findIndexOfClosestPoint(points []Point, x int, y int) int {
	distances := make(map[int]int)
	for i := 0; i < len(points); i++ {
		distances[i] = manhattanDist(points[i].x, points[i].y, x, y)
	}

	sortedList := datastructure.SortMapByValue(distances, false)

	if sortedList[0].Value == sortedList[1].Value {
		return -1
	} else {
		return sortedList[0].Key
	}
}

func manhattanDist(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs((float64)(x1-x2))) + int(math.Abs((float64)(y1-y2)))
}

func getIndexesOfBounds(points []Point, minX int, minY int, maxX int, maxY int) []int {
	var indexes []int
	for i := 0; i < len(points); i++ {
		if points[i].x == minX || points[i].x == maxX || points[i].y == minY || points[i].y == maxY {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func findBounds(points []Point) (int, int, int, int) {
	minX := points[0].x
	minY := points[0].y
	maxX := points[0].x
	maxY := points[0].y
	for i := 0; i < len(points); i++ {
		if points[i].x < minX {
			minX = points[i].x
		}
		if points[i].x > maxX {
			maxX = points[i].x
		}
		if points[i].y < minY {
			minY = points[i].y
		}
		if points[i].y > maxY {
			maxY = points[i].y
		}
	}
	return minX, minY, maxX, maxY
}

func buildPointList(lines []string) []Point {
	var points []Point
	for i := 0; i < len(lines); i++ {
		points = append(points, buildPointFromLine(lines[i]))
	}
	return points
}

func buildPointFromLine(line string) Point {
	temp := strings.Split(line, ",")
	x, _ := strconv.Atoi(strings.Trim(temp[0], " "))
	y, _ := strconv.Atoi(strings.Trim(temp[1], " "))
	return Point{x: x, y: y}
}
