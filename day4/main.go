package main

import (
	"strings"
	"github.com/cfagiani/aoc2018/util"
	"sort"
	"fmt"
	"strconv"
)

type Guard struct {
	id     int
	shifts map[string][60]int
}

func main() {
	inputString := util.ReadFileAsString("input/day4.input")
	lines := strings.Split(inputString, "\n")
	guards := populateGuardShifts(lines)
	id, minute := part1(guards)
	fmt.Printf("Product of max is %d\n", (id * minute))
	id, minute = part2(guards)
	fmt.Printf("Product of max is %d\n", (id * minute))
}

//Returns the id of the guard and minute that is most likely to be sleeping
func part1(guards map[int]Guard) (int, int) {
	maxVal := 0
	maxId := -1
	for _, v := range guards {

		totalSleep := 0
		for _, shift := range v.shifts {
			for _, d := range shift {
				totalSleep += d

			}
		}
		if totalSleep > maxVal {
			maxVal = totalSleep
			maxId = v.id
		}
	}
	//now find minute most likely to be asleep
	minMap := [60]int{}
	for _, shift := range guards[maxId].shifts {
		for i, v := range shift {
			minMap[i] += v
		}
	}
	maxIdx := 0

	for i, v := range minMap {
		if v >= minMap[maxIdx] {
			maxIdx = i
		}
	}
	return maxId, maxIdx
}

//Returns the id and minute of the guard that is most often alseep on the same minute
func part2(guards map[int]Guard) (int, int) {
	minMaps := make(map[int][60]int)
	for id, guard := range guards {
		minMap := [60]int{}
		for _, shift := range guard.shifts {
			for i, v := range shift {
				minMap[i] += v
			}
		}
		minMaps[id] = minMap
	}

	//now find the max
	maxId := -1
	maxVal := 0
	maxMin := -1
	for id, minMap := range minMaps {
		for i, v := range minMap {
			if v >= maxVal {
				maxId = id
				maxVal = v
				maxMin = i
			}
		}
	}
	return maxId, maxMin
}

func populateGuardShifts(lines []string) map[int]Guard {
	guards := make(map[int]Guard)
	sort.Strings(lines)
	var curGuard = Guard{id: -1}
	prevTime := 0
	curDay := ""
	isSleeping := 0
	for i := 0; i < len(lines); i++ {
		day, hour, minute := getDayTimeFromLine(lines[i])
		if strings.Contains(lines[i], "begins shift") {
			if curGuard.id != -1 {
				arr := curGuard.shifts[curDay]
				for j := prevTime; j < 60; j++ {
					arr[j] = isSleeping
				}
				curGuard.shifts[curDay] = arr
			}
			y, m, d := splitDate(day)
			if hour > 0 {
				d++
			}
			dayKey := fmt.Sprintf("%d-%d-%d", y, m, d)
			curDay = dayKey
			isSleeping = 0
			prevTime = minute
			guardId := getGuardId(lines[i])
			existingGuard, present := guards[guardId]
			if present {
				curGuard = existingGuard
			} else {
				curGuard = Guard{id: guardId, shifts: make(map[string][60]int)}
				guards[guardId] = curGuard
			}
			curGuard.shifts[dayKey] = [60]int{}
		} else {
			arr := curGuard.shifts[curDay]
			for j := prevTime; j < minute; j++ {
				arr[j] = isSleeping
			}
			curGuard.shifts[curDay] = arr
			if strings.Contains(lines[i], "falls asleep") {
				isSleeping = 1
			} else if strings.Contains(lines[i], "wakes up") {
				isSleeping = 0
			}
			prevTime = minute
		}
	}
	//handle last one
	arr := curGuard.shifts[curDay]
	for j := prevTime; j < 60; j++ {
		arr[j] = isSleeping
	}
	curGuard.shifts[curDay] = arr
	return guards
}

func getGuardId(line string) int {
	parts := strings.Split(line, "#")
	id, _ := strconv.Atoi(strings.Split(parts[1], " ")[0])
	return id
}

func getDayTimeFromLine(line string) (string, int, int) {
	parts := strings.Split(line, " ")
	day := strings.Replace(parts[0], "[", "", -1)
	timeString := strings.Replace(parts[1], "]", "", -1)
	timeParts := strings.Split(timeString, ":")
	hour, _ := strconv.Atoi(strings.Trim(timeParts[0], " "))
	minute, _ := strconv.Atoi(strings.Trim(timeParts[1], " "))
	return day, hour, minute
}

func splitDate(dateString string) (int, int, int) {
	parts := strings.Split(dateString, "-")
	y, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	d, _ := strconv.Atoi(parts[2])
	return y, m, d
}
