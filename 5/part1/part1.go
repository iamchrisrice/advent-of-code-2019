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

type Flags struct {
	p1 bool
	p2 bool
	p3 bool
}

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
	return Instruction{opcode, Flags{p1mode == 1, p2mode == 1, p3mode == 1}}
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
		case 99:
			return
		}
	}
}

func add(program []int, position *int, flags Flags) {
	var l1, l2, l3 int
	if flags.p1 {
		l1 = *position + 1
	} else {
		l1 = program[*position+1]
	}
	if flags.p2 {
		l2 = *position + 2
	} else {
		l2 = program[*position+2]
	}
	l3 = program[*position+3]
	program[l3] = program[l1] + program[l2]
	*position = *position + 4
}

func multiply(program []int, position *int, flags Flags) {
	var l1, l2, l3 int
	if flags.p1 {
		l1 = *position + 1
	} else {
		l1 = program[*position+1]
	}
	if flags.p2 {
		l2 = *position + 2
	} else {
		l2 = program[*position+2]
	}
	l3 = program[*position+3]
	program[l3] = program[l1] * program[l2]
	*position = *position + 4
}

func input(program []int, position *int, flags Flags) {
	location := program[*position+1]
	var i int
	fmt.Print("> ")
	fmt.Scan(&i)
	program[location] = i
	*position = *position + 2
}

func output(program []int, position *int, flags Flags) {
	var l1 int
	if flags.p1 {
		l1 = *position + 1
	} else {
		l1 = program[*position+1]
	}
	fmt.Println(program[l1])
	*position = *position + 2
}

func main() {
	const filename = "../input.txt"
	buf, _ := ioutil.ReadFile(filename)
	input := string(buf)
	sli := toIntSlice(input)
	emulate(sli)
}
