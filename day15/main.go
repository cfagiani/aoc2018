package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/datastructure"
	"github.com/cfagiani/aoc2018/util"
	"math"
	"sort"
	"strings"
)

const StartingPower = 3
const StartingHitPoints = 200

type Unit struct {
	Pos       datastructure.IntPair
	Type      byte
	Power     int
	HitPoints int
}

type PossibleMove struct {
	currentLoc  datastructure.IntPair
	currentDist int
	startMove   datastructure.IntPair
}

type Path struct {
	steps []datastructure.IntPair
}

type UnitList []*Unit

func main() {
	inputString := util.ReadFileAsString("input/day15.input")
	state, units := initialize(inputString)
	fmt.Printf("Initial State\n")
	util.PrintByteArray(state)
	part1(state, units)
	part2(inputString)
}

func initialize(input string) ([][]byte, []*Unit) {
	lines := strings.Split(input, "\n")
	state := make([][]byte, len(lines))
	var units []*Unit
	for i := 0; i < len(lines); i++ {
		state[i] = make([]byte, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			state[i][j] = lines[i][j]
			if state[i][j] == 'E' || state[i][j] == 'G' {
				units = append(units, &Unit{
					Pos:       datastructure.IntPair{A: j, B: i},
					Type:      state[i][j],
					Power:     StartingPower,
					HitPoints: StartingHitPoints})
			}
		}
	}
	return state, units
}

func part1(state [][]byte, units []*Unit) {
	rounds := 0
	var endedEarly bool
	for done := false; !done; {
		state, endedEarly = performRound(state, units)
		if !endedEarly {
			rounds++
		}
		done = len(findEnemies('G', units)) == 0 || len(findEnemies('E', units)) == 0
		fmt.Printf("After %d, sum is: %d\n", rounds, getSumOfPoints(units))
		util.PrintByteArray(state)
	}

	fmt.Printf("Results of %d rounds of combat: %d\n", rounds, getSumOfPoints(units)*rounds)
}

func part2(inputString string) {
	for power := StartingPower; ; power++ {
		state, units := initialize(inputString)
		//adjust power for all elves
		elves := findEnemies('G', units)
		for i := 0; i < len(elves); i++ {
			elves[i].Power = power
		}
		//get initial count of elves
		elfCount := len(findEnemies('G', units))
		var endedEarly bool
		rounds := 0
		for done := false; !done; {
			state, endedEarly = performRound(state, units)
			done = len(findEnemies('G', units)) < elfCount || len(findEnemies('E', units)) == 0
			if !endedEarly {
				rounds++
			}
		}
		if len(findEnemies('G', units)) == elfCount {
			fmt.Printf("Min power for no elf losses is %d\n", power)
			fmt.Printf("Results of %d rounds of combat: %d\n", rounds, getSumOfPoints(units)*rounds)
			break
		}
	}
}

func getSumOfPoints(units []*Unit) int {
	sum := 0
	for i := 0; i < len(units); i++ {
		if units[i].HitPoints > 0 {
			sum += units[i].HitPoints
		}
	}
	return sum
}

func performRound(state [][]byte, units []*Unit) ([][]byte, bool) {
	sort.Sort(UnitList(units))
	endedEarly := false
	for i := 0; i < len(units); i++ {
		if units[i].HitPoints <= 0 {
			//if a unit is dead, it can't move
			continue
		}
		if len(findEnemies(units[i].Type, units)) == 0 {
			endedEarly = true
			break
		}
		newDest, isRealMove := getMove(*units[i], state)
		if isRealMove {
			state = move(units[i], newDest, state)
		}
		state = attackIfPossible(units[i], units, state)
	}

	return state, endedEarly
}

//BFS for moves
func getMove(unit Unit, state [][]byte) (datastructure.IntPair, bool) {
	queue := datastructure.NewPriorityQueue(5)
	visited := make([][]bool, len(state))
	for i := 0; i < len(state); i++ {
		visited[i] = make([]bool, len(state[i]))
	}
	visited[unit.Pos.B][unit.Pos.A] = true // mark starting loc as visited
	neighborOffsets := []datastructure.IntPair{{A: 0, B: -1}, {A: 0, B: 1}, {A: -1, B: 0}, {A: 1, B: 0}}
	//initialize with neighbors, exiting if we're already next to an enemy
	for i := 0; i < len(neighborOffsets); i++ {
		loc := datastructure.IntPair{A: unit.Pos.A + neighborOffsets[i].A, B: unit.Pos.B + neighborOffsets[i].B}
		if state[loc.B][loc.A] == '.' {
			queue.Enqueue(PossibleMove{loc, 1, loc})
			visited[loc.B][loc.A] = true
		} else if isEnemyAtLoc(unit, loc, state) {
			return datastructure.IntPair{A: -1, B: -1}, false
		}
	}
	for queue.Len() > 0 {
		candidate := queue.Dequeue().(PossibleMove)
		if isEnemyAtLoc(unit, candidate.currentLoc, state) {
			return candidate.startMove, true
		}
		for i := 0; i < len(neighborOffsets); i++ {
			loc := datastructure.IntPair{A: candidate.currentLoc.A + neighborOffsets[i].A, B: candidate.currentLoc.B + neighborOffsets[i].B}
			if (state[loc.B][loc.A] == '.' || isEnemyAtLoc(unit, loc, state)) && !visited[loc.B][loc.A] {
				queue.Enqueue(PossibleMove{loc, candidate.currentDist + 1, candidate.startMove})
				visited[loc.B][loc.A] = true
			}
		}
	}
	return datastructure.IntPair{A: -1, B: -1}, false
}

func attackIfPossible(unit *Unit, allUnits []*Unit, state [][]byte) [][]byte {
	candidates := findEnemies(unit.Type, allUnits)
	var selectedUnit *Unit
	for i := 0; i < len(candidates); i++ {
		//check if adjacent
		if (unit.Pos.A == candidates[i].Pos.A && math.Abs((float64)(unit.Pos.B-candidates[i].Pos.B)) == 1.0) ||
			(unit.Pos.B == candidates[i].Pos.B && math.Abs((float64)(unit.Pos.A-candidates[i].Pos.A)) == 1.0) {
			if selectedUnit == nil || candidates[i].HitPoints < selectedUnit.HitPoints {
				selectedUnit = candidates[i]
			} else if selectedUnit != nil && candidates[i].HitPoints == selectedUnit.HitPoints {
				if candidates[i].Pos.B < selectedUnit.Pos.B ||
					(candidates[i].Pos.B == selectedUnit.Pos.B && candidates[i].Pos.A < selectedUnit.Pos.A) {
					selectedUnit = candidates[i]
				}
			}
		}
	}
	if selectedUnit != nil {
		selectedUnit.HitPoints -= unit.Power
		if selectedUnit.HitPoints <= 0 {
			state[selectedUnit.Pos.B][selectedUnit.Pos.A] = '.'
		}
	}
	return state
}

func move(unit *Unit, dest datastructure.IntPair, state [][]byte) [][]byte {
	state[unit.Pos.B][unit.Pos.A] = '.'
	unit.Pos.A = dest.A
	unit.Pos.B = dest.B
	state[unit.Pos.B][unit.Pos.A] = unit.Type
	return state
}



func findEnemies(unitType byte, units []*Unit) []*Unit {
	var enemies []*Unit
	for i := 0; i < len(units); i++ {
		if units[i].HitPoints > 0 && units[i].Type != unitType {
			enemies = append(enemies, units[i])
		}
	}
	return enemies
}

func isEnemyAtLoc(unit Unit, loc datastructure.IntPair, state [][]byte) bool {
	return state[loc.B][loc.A] != '.' && state[loc.B][loc.A] != '#' && state[loc.B][loc.A] != unit.Type
}

func (a UnitList) Len() int { return len(a) }
func (a UnitList) Less(i, j int) bool {
	return a[i].Pos.Less(a[j].Pos, false)
}

func (a UnitList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a PossibleMove) Compare(b datastructure.Comparable) int {
	other := b.(PossibleMove)
	if a.currentDist < other.currentDist {
		return -1
	} else if a.currentDist > other.currentDist {
		return 1
	}
	if a.currentLoc.Equals(other.currentLoc) {
		if a.startMove.Less(other.startMove, false) {
			return -1
		} else {
			return 1
		}
	}
	if a.currentLoc.Less(other.currentLoc, false) {
		return -1
	}
	return 1
}
