package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	opcode int
	flags  Flags
}

type Flags = map[int]bool

type Params = map[int]int

func toIntSlice(input string) []int {
	var sli []int
	for _, number := range strings.Split(input, ",") {
		integer, _ := strconv.Atoi(number)
		sli = append(sli, integer)
	}
	return sli
}

func toInstruction(input int) Instruction {
	inst := strconv.Itoa(input)
	for len(inst) < 5 {
		inst = "0" + inst
	}
	opcode, _ := strconv.Atoi(string(inst[3]) + string(inst[4]))
	p1mode, _ := strconv.Atoi(string(inst[2]))
	p2mode, _ := strconv.Atoi(string(inst[1]))
	p3mode, _ := strconv.Atoi(string(inst[0]))
	return Instruction{opcode, Flags{1: p1mode == 1, 2: p2mode == 1, 3: p3mode == 1}}
}

func getParams(paramCount int, program []int, position *int, flags Flags) Params {
	params := make(Params)
	for count := 1; count <= paramCount; count++ {
		if flags[count] {
			params[count] = *position + count
		} else {
			params[count] = program[*position+count]
		}
	}
	return params
}

func emulate(program []int) {
	position := 0
	for {
		instruction := toInstruction(program[position])

		switch instruction.opcode {
		case 1:
			add(program, &position, instruction.flags)
		case 2:
			multiply(program, &position, instruction.flags)
		case 3:
			input(program, &position, instruction.flags)
		case 4:
			output(program, &position, instruction.flags)
		case 5:
			jumpIfTrue(program, &position, instruction.flags)
		case 6:
			jumpIfFalse(program, &position, instruction.flags)
		case 7:
			lessThan(program, &position, instruction.flags)
		case 8:
			equals(program, &position, instruction.flags)
		case 99:
			return
		}
	}
}

func add(program []int, position *int, flags Flags) {
	flags[3] = false
	params := getParams(3, program, position, flags)
	program[params[3]] = program[params[1]] + program[params[2]]
	*position = *position + 4
}

func multiply(program []int, position *int, flags Flags) {
	flags[3] = false
	params := getParams(3, program, position, flags)
	program[params[3]] = program[params[1]] * program[params[2]]
	*position = *position + 4
}

func input(program []int, position *int, flags Flags) {
	flags[1] = false
	params := getParams(1, program, position, flags)
	var i int
	fmt.Print("> ")
	fmt.Scan(&i)
	program[params[1]] = i
	*position = *position + 2
}

func output(program []int, position *int, flags Flags) {
	params := getParams(1, program, position, flags)
	fmt.Println(program[params[1]])
	*position = *position + 2
}

func jumpIfTrue(program []int, position *int, flags Flags) {
	params := getParams(2, program, position, flags)
	if program[params[1]] != 0 {
		*position = program[params[2]]
	} else {
		*position = *position + 3
	}
}

func jumpIfFalse(program []int, position *int, flags Flags) {
	params := getParams(2, program, position, flags)
	if program[params[1]] == 0 {
		*position = program[params[2]]
	} else {
		*position = *position + 3
	}
}

func lessThan(program []int, position *int, flags Flags) {
	flags[3] = false
	params := getParams(3, program, position, flags)
	if program[params[1]] < program[params[2]] {
		program[params[3]] = 1
	} else {
		program[params[3]] = 0
	}
	*position = *position + 4
}

func equals(program []int, position *int, flags Flags) {
	flags[3] = false
	params := getParams(3, program, position, flags)
	if program[params[1]] == program[params[2]] {
		program[params[3]] = 1
	} else {
		program[params[3]] = 0
	}
	*position = *position + 4
}

func main() {
	const filename = "../input.txt"
	buf, _ := ioutil.ReadFile(filename)
	input := string(buf)
	sli := toIntSlice(input)
	emulate(sli)
}
