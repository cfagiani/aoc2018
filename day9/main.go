package main

import (
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"strings"
	"fmt"
	"math"
)

type Marble struct {
	val  int
	next *Marble
	prev *Marble
}

func main() {
	inputString := util.ReadFileAsString("input/day9.input")
	parts := strings.Split(inputString, " ")
	players, _ := strconv.Atoi(parts[0])
	lastMarble, _ := strconv.Atoi(parts[6])
	part1(players, lastMarble)
	part2(players, lastMarble)
}

func part2(players int, lastMarble int) {
	max := findMaxScore(players, lastMarble*100)
	fmt.Printf("Winning score is %d\n", max)
}

func part1(players int, lastMarble int) {
	max := findMaxScore(players, lastMarble)
	fmt.Printf("Winning score is %d\n", max)
}

func findMaxScore(players int, lastMarble int) int {
	scores := playGame(players, lastMarble)
	max := -1
	for i := 0; i < len(scores); i++ {
		if scores[i] > max {
			max = scores[i]
		}
	}
	return max
}

func playGame(players int, lastMarble int) []int {
	scores := []int{}
	for i := 0; i < players; i++ {
		scores = append(scores, 0)
	}

	curMarble := &Marble{val: 0}
	curMarble.next = curMarble
	curMarble.prev = curMarble
	for i := 1; i <= lastMarble; i++ {
		if i > 0 && i%23 == 0 {
			nextCur, curScore := removeMarble(curMarble, -7)
			curScore += i
			scores[(i-1)%players] += curScore
			curMarble = nextCur
		} else {
			curMarble = insertMarble(curMarble, i)
		}
	}
	return scores
}

func advanceCur(curMarble *Marble, offset int) *Marble {
	targetPos := curMarble
	for i := 0; i < int(math.Abs(float64(offset))); i++ {
		if offset > 0 {
			targetPos = targetPos.next
		} else {
			targetPos = targetPos.prev
		}
	}
	return targetPos
}

func removeMarble(curMarble *Marble, offset int) (*Marble, int) {
	targetPos := advanceCur(curMarble, offset)
	targetPos.prev.next = targetPos.next
	targetPos.next.prev = targetPos.prev
	return targetPos.next, targetPos.val

}

func insertMarble(curMarble *Marble, val int) *Marble {
	insertBefore := advanceCur(curMarble, 2)
	newMarble := &Marble{val: val, next: insertBefore, prev: insertBefore.prev}
	insertBefore.prev.next = newMarble

	insertBefore.prev = newMarble

	return newMarble
}
