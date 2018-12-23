package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"strings"
)

func main() {
	inputString := util.ReadFileAsString("input/day14.input")
	scores := []byte{'3', '7'}
	recipeCount, _ := strconv.Atoi(inputString)
	part1(scores, recipeCount)
	scores = []byte{'3', '7'}
	part2(scores, recipeCount)
}

func part1(scores []byte, recipeCount int) {
	curPointer := []int{0, 1}
	for len(scores) < recipeCount {
		scores, curPointer = makeNextRecipe(scores, curPointer)
	}
	//now do 10 more
	startIdx := len(scores)
	for i := 0; i < 10; i++ {
		scores, curPointer = makeNextRecipe(scores, curPointer)
	}
	fmt.Printf("The last 10 scores are %s\n", scores[startIdx:startIdx+10])
}

func part2(scores []byte, searchValue int) {
	curPointer := []int{0, 1}
	searchString := strconv.Itoa(searchValue)
	for {
		scores, curPointer = makeNextRecipe(scores, curPointer)
		//only check every 100000 recipes so we don't go super slow
		if len(scores)%100000 == 0 {
			idx := strings.Index(string(scores), searchString)
			if idx >= 0 {
				fmt.Printf("Search string appears after %d scores\n", idx)
				break
			}
		}
	}
}

func makeNextRecipe(scores []byte, curPointers []int) ([]byte, []int) {
	newScore := []byte(strconv.Itoa(int(scores[curPointers[0]] - '0' + scores[curPointers[1]] - '0')))
	scores = append(scores, newScore...)
	for i := 0; i < len(curPointers); i++ {
		curPointers[i] = (curPointers[i] + 1 + int(scores[curPointers[i]]-'0')) % len(scores)
	}
	return scores, curPointers
}
