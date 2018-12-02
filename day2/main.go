package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strings"
)

func main() {
	inputString := util.ReadFileAsString("input/day2.input")
	lines := strings.Split(inputString, "\n")
	part1(lines)
	fmt.Printf("Match %s\n", getMatchingStrings(part2(lines)))
}

func part2(lines []string) (string, string) {

	for i := 1; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			diffCount := 0
			for k := 0; k < len(lines[i]); k++ {
				if lines[i][k] != lines[j][k] {
					diffCount++
				}
			}
			if diffCount == 1 {
				return lines[i], lines[j]
			}
		}

	}
	return "", ""
}

func part1(lines []string) {
	hasTwo := 0
	hasThree := 0
	for i := 0; i < len(lines); i++ {
		histogram := computeHistogram(lines[i])
		foundTwo := false
		foundThree := false
		for _, v := range histogram {
			if v == 2 && !foundTwo {
				hasTwo++
				foundTwo = true
			}
			if v == 3 && !foundThree {
				hasThree++
				foundThree = true
			}
		}
	}

	fmt.Printf("Checksum  is %d\n", hasTwo*hasThree)
}

func computeHistogram(line string) map[uint8]int {
	histo := make(map[uint8]int)
	for i := 0; i < len(line); i++ {
		curCount, found := histo[line[i]]
		if found {
			histo[line[i]] = curCount + 1
		} else {
			histo[line[i]] = 1
		}
	}
	return histo
}

func getMatchingStrings(a string, b string) string {
	outString := ""
	for pos, char := range a {
		if a[pos] == b[pos] {
			outString = fmt.Sprintf("%s%c", outString, char)
		}
	}

	return outString
}
