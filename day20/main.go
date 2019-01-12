package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
)

type Room struct {
	X    int
	Y    int
	Dist int
	N    *Room
	S    *Room
	E    *Room
	W    *Room
}

var Unknown = &Room{}

var roomMap = make(map[string]*Room)

func main() {
	inputString := util.ReadFileAsString("input/day20.input")
	startPoint := makeRoom()
	buildMap(inputString, startPoint)
	part1()
	part2()
}

func part1() {
	//just find the room with greatest distance
	maxDist := 0
	for _, val := range roomMap {
		if val.Dist > maxDist {
			maxDist = val.Dist
		}
	}
	fmt.Printf("Max distance is: %d doors\n", maxDist)
}

func part2() {
	count := 0
	threshold := 1000
	for _, val := range roomMap {
		if val.Dist >= threshold {
			count++
		}
	}
	fmt.Printf("%d rooms are at least %d doors away\n", count, threshold)
}

func buildMap(input string, startPoint *Room) {
	curRoom := startPoint
	for i := 0; i < len(input); i++ {
		if input[i] == '^' {
			continue
		} else if input[i] == '$' {
			return
		} else if input[i] == '(' {
			options, lastIdx := getOptionsStrings(input[i:])
			for j := 0; j < len(options); j++ {
				buildMap(options[j], curRoom)
			}
			i += lastIdx
		} else {
			curRoom = addNeighbor(curRoom, input[i])
		}
	}
}

func getOptionsStrings(val string) ([]string, int) {
	openCount := 0
	var idx int
	var options []string
	curOption := ""
	for idx = 0; idx < len(val); idx++ {
		if val[idx] == '(' {
			openCount++
		} else if val[idx] == ')' {
			openCount--
		}

		if openCount == 0 {
			//we're done
			break
		} else if openCount == 1 {
			//if we're only 1 level deep in the parens, then this is an option
			if val[idx] == '|' {
				options = append(options, curOption)
				curOption = ""
				continue
			}
		}
		//if we're here, we just need to accumulate the current character as long as it's not the open/close paren for level 1
		if openCount > 1 || (val[idx] != '(' && val[idx] != ')') {
			curOption += string(val[idx])
		}
	}
	options = append(options, curOption)
	return options, idx
}

func makeRoom() *Room {
	return &Room{X: 0, Y: 0, Dist: 0, N: Unknown, S: Unknown, E: Unknown, W: Unknown}
}

func addNeighbor(room *Room, dir byte) *Room {
	x := room.X
	y := room.Y
	switch dir {
	case 'N':
		y--
	case 'S':
		y++
	case 'E':
		x++
	case 'W':
		x--
	}
	key := fmt.Sprintf("%d,%d", x, y)
	neighbor, ok := roomMap[key]
	if ok {
		//did we find a shorter path to this room?
		if room.Dist+1 < neighbor.Dist {
			neighbor.Dist = room.Dist + 1
		}
	} else {
		neighbor = makeRoom()
		neighbor.Dist = room.Dist + 1
		neighbor.X = x
		neighbor.Y = y
		roomMap[key] = neighbor
	}
	switch dir {
	case 'N':
		room.N = neighbor
	case 'S':
		room.S = neighbor
	case 'E':
		room.E = neighbor
	case 'W':
		room.W = neighbor
	}
	return neighbor
}
