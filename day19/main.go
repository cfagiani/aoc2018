package main

import (
	"fmt"
	"github.com/cfagiani/aoc2018/util"
	"strconv"
	"strings"
)

type Op func(int, int, int, []int)

var opcodeFuncs = map[string]Op{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr}

func main() {
	inputString := util.ReadFileAsString("input/day19.input")
	lines := strings.Split(inputString, "\n")
	ipRegister, _ := strconv.Atoi(strings.Split(lines[0], " ")[1])
	program := lines[1:]
	part1(program, ipRegister)
	part2(program, ipRegister)
}

func part1(program []string, ipRegister int) {
	registers := make([]int, 6)
	runProgram(registers, program, ipRegister)
	fmt.Printf("After termination, register 0 contains %d\n", registers[0])
}

func part2(program []string, ipRegister int) {
	registers := make([]int, 6)
	registers[0] = 1
	//from analysis of program it looks like it is summing the factors of 10551432 (r2)
	runProgram(registers, program, ipRegister)
	fmt.Printf("After termination, register 0 contains %d\n", registers[0])

}

func translateProgram(program []string, ipRegister int) []string {

	translated := make([]string, len(program))
	for i := 0; i < len(program); i++ {
		parts := strings.Split(program[i], " ")
		opCode := parts[0]
		args := stringArrayToIntArray(parts[1:])
		if args[2] == ipRegister {
			if opCode == "seti" {
				translated[i] = fmt.Sprintf("GOTO %d", args[0])
			} else if opCode == "addi" {
				translated[i] = fmt.Sprintf("GOTO %d", args[1]+i)
			} else if opCode == "addr" || opCode == "mulr" {
				op := "+"
				if opCode == "mulr" {
					op = "*"
				}
				translated[i] = fmt.Sprintf("GOTO %s %s %s", translateRegister(i+1, ipRegister, args[0]), op, translateRegister(i+1, ipRegister, args[1]))
			}
			continue
		}

		if opCode == "addi" || opCode == "muli" {
			op := "+"
			if opCode == "muli" {
				op = "*"
			}
			translated[i] = fmt.Sprintf("%s = %s %s %d", translateRegister(i, ipRegister, args[2]), translateRegister(i, ipRegister, args[1]), op, args[1])
		} else if opCode == "addr" || opCode == "mulr" {
			op := "+"
			if opCode == "mulr" {
				op = "*"
			}
			translated[i] = fmt.Sprintf("%s = %s %s %s", translateRegister(i, ipRegister, args[2]), translateRegister(i, ipRegister, args[0]), op, translateRegister(i, ipRegister, args[1]))
		} else if opCode == "seti" {
			translated[i] = fmt.Sprintf("%s = %d", translateRegister(i, ipRegister, args[2]), args[0])
		} else if opCode == "setr" {
			translated[i] = fmt.Sprintf("%s = %s", translateRegister(i, ipRegister, args[2]), translateRegister(i, ipRegister, args[0]))
		} else if opCode == "eqrr" || opCode == "gtrr" {
			op := "=="
			if opCode == "gtrr" {
				op = ">"
			}
			translated[i] = fmt.Sprintf("if %s %s %s then %s = 1 else %s = 0", translateRegister(i, ipRegister, args[0]), op, translateRegister(i, ipRegister, args[1]), translateRegister(i, ipRegister, args[2]), translateRegister(i, ipRegister, args[2]))
		}
	}

	return translated
}

func translateRegister(ip int, ipRegister int, regNum int) string {
	if regNum == ipRegister {
		return fmt.Sprintf("%d", ip)
	} else {
		return string(65 + regNum)
	}

}

func runProgram(registers []int, program []string, ipRegister int) {
	for ip := registers[ipRegister]; ip < len(program); ip++ {
		if ip < 0 {
			break
		} else if ip == 3 && registers[1] != 0 {
			//based on analysis of translated program, this performs the work of the main loop that occurs between instructions 2 and 13
			for ; registers[1] <= registers[2]; registers[1]++ {
				if registers[2]%registers[1] == 0 {
					registers[0] += registers[1]
				}
			}
			ip = 12
			return
			//continue
		}

		registers[ipRegister] = ip
		parts := strings.Split(program[ip], " ")
		opCode := parts[0]
		args := stringArrayToIntArray(parts[1:])
		opcodeFuncs[opCode](args[0], args[1], args[2], registers)
		if args[2] == ipRegister {
			ip = registers[ipRegister]
		}
	}
}

func addr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] + registers[b]
}

func addi(a int, b int, c int, registers []int) {
	registers[c] = registers[a] + b
}

func mulr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] * registers[b]
}

func muli(a int, b int, c int, registers []int) {
	registers[c] = registers[a] * b
}

func banr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] & registers[b]
}

func bani(a int, b int, c int, registers []int) {
	registers[c] = registers[a] & b
}

func borr(a int, b int, c int, registers []int) {
	registers[c] = registers[a] | registers[b]
}

func bori(a int, b int, c int, registers []int) {
	registers[c] = registers[a] | b
}

func setr(a int, b int, c int, registers []int) {
	registers[c] = registers[a]
}

func seti(a int, b int, c int, registers []int) {
	registers[c] = a
}

func gtir(a int, b int, c int, registers []int) {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtri(a int, b int, c int, registers []int) {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func gtrr(a int, b int, c int, registers []int) {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqir(a int, b int, c int, registers []int) {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqri(a int, b int, c int, registers []int) {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func eqrr(a int, b int, c int, registers []int) {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
}

func stringArrayToIntArray(input []string) []int {
	results := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		results[i], _ = strconv.Atoi(input[i])
	}
	return results
}
