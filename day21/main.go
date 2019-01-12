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
	inputString := util.ReadFileAsString("input/day21.input")
	lines := strings.Split(inputString, "\n")
	ipRegister, _ := strconv.Atoi(strings.Split(lines[0], " ")[1])
	program := lines[1:]
	part1(program, ipRegister)
}

func part1(program []string, ipRegister int) {

	registers := make([]int, 6)
	registers[0] = 16134795 // This number obtained by doing a hand-analysis of the input program and then using the debugger to inspect register 0 when we first hit the eqrr instruction
	count := runProgram(registers, program, ipRegister)

	fmt.Printf("Ran %d instructions\n", count)
	runTranslatedProgram(16134795)
}

func runProgram(registers []int, program []string, ipRegister int) int {
	counter := 0
	for ip := registers[ipRegister]; ip < len(program); ip++ {
		if ip < 0 {
			break
		}
		counter++
		registers[ipRegister] = ip
		parts := strings.Split(program[ip], " ")
		opCode := parts[0]
		args := stringArrayToIntArray(parts[1:])
		opcodeFuncs[opCode](args[0], args[1], args[2], registers)
		if args[2] == ipRegister {
			ip = registers[ipRegister]
		}
	}
	return counter
}

func runTranslatedProgram(val int) {

	r0, r1, r3, r4, r5 := val, 0, 0, 0, 0
LBL0:
	r3 = 123      //seti 123 0 3
	r3 = r3 & 456 // bani 3 456 3
	if r3 == 72 { // eqri 3 72 3
		r3 = 1
		goto LBL5
	} else {
		r3 = 0
		goto LBL0
	}
	//LBL3: addr 3 2 2 // SEE ABOVE
	//LBL4: seti 0 0 2 // SEE ABOVE
LBL5:
	r3 = 0             // seti 0 4 3
	r4 = 65536 | r3    // bori 3 65536 4
LBL7:
	r3 = 1107552       // seti 1107552 3 3
	r5 = r4 | 255      // bani 4 255 5
	r3 = r3 + r5       // addr 3 5 3
	r3 = 16777215 | r3 // bani 3 16777215 3
	r3 = r3 * 65899    // muli 3 65899 3
	r3 = r3 & 16777215 // bani 3 16777215 3
	if 256 > r4 { //gtir 256 4 5
		r5 = 1
		goto LBL27
	} else {
		r5 = 0
		goto LBL17
	}
	//LBL14: addr 5 2 2 //SEE ABOVE
	//LBL15: addi 2 1 2 //SEE ABOVE
	//LBL16: seti 27 0 2   // SEE ABOVE
	r5 = 0        // seti 0 2 5
LBL17:
	r1 = r5 + 1   // addi 5 1 1
	r1 = r1 * 256 // muli 1 256 1
	if r1 > r4 { // gtrr 1 4 1
		r1 = 1
		goto LBL25
	} else {
		r1 = 0
		goto LBL24
	}
	//LBL21: addr 1 2 2 // SEE ABOVE
	//LBL22: addi 2 1 2 // SEE ABOVE
	//LBL23: seti 25 3 2  // SEE ABOVE
LBL24:
	r5 = r5 + 1 // addi 5 1 5
LBL25:
	goto LBL17  //seti 17 3 2
	r4 = r5     // setr 5 3 4
LBL27:
	goto LBL7   // seti 7 4 2
	if r3 == r0 { // eqrr 3 0 5
		r5 = 1
		return
	} else {
		r5 = 0
		goto LBL5
	}
	//LBL29: addr 5 2 2 // SEE ABOVE
	//LBL30: seti 5 8 2 // SEE ABOVE

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
