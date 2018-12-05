package main

import (
	"github.com/cfagiani/aoc2018/util"
	"math"
	"fmt"
	"strings"
)

func main() {
	inputString := util.ReadFileAsString("input/day5.input")

	fmt.Printf("Original length: %d\n", len(inputString))
	part1(inputString)
	part2(inputString)
}

func part2(inputString string) {
	minLen := len(inputString)
	for i := 65; i < 91; i++ {
		poly := processPolymer(inputString, fmt.Sprintf("%s%s", string(i), string(i+32)))
		if len(poly) < minLen {
			minLen = len(poly)
		}
	}
	fmt.Printf("Min Length polymer is %d\n", minLen)
}

func part1(inputString string) {
	polymer := processPolymer(inputString, "")
	fmt.Printf("Final length: %d\n", len(polymer))
}

func processPolymer(input string, replacements string) string {
	temp := input
	for i := range replacements {
		temp = strings.Replace(temp, string(replacements[i]), "", -1)
	}
	curString := []byte(temp)
	for i := 0; i < len(curString); i++ {
		if i < len(curString)-1 {
			if int(math.Abs(float64(curString[i])-float64(curString[i+1]))) == 32 {
				curString = append(curString[:i], curString[i+2:]...)
				i = -1 // need to set to -1 since the for loop is going to increment
			}
		}
	}
	return string(curString)
}
