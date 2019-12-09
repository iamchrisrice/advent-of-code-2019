package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func toIntSlice(input string) []int {
	var sli []int
	for _, number := range strings.Split(input, ",") {
		integer, _ := strconv.Atoi(number)
		sli = append(sli, integer)
	}
	return sli
}

func emulate(program []int) {
	position := 0
	for {
		switch program[position] {
		case 1:
			location1 := program[position+1]
			location2 := program[position+2]
			location3 := program[position+3]
			program[location3] = program[location1] + program[location2]
		case 2:
			location1 := program[position+1]
			location2 := program[position+2]
			location3 := program[position+3]
			program[location3] = program[location1] * program[location2]
		case 99:
			return
		}
		position = position + 4
	}
}

func checkForDesiredOutput(program []int, output int) (int, int) {
	var noun, verb int
Loop:
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			memory := make([]int, len(program))
			copy(memory, program)
			memory[1] = i
			memory[2] = j
			emulate(memory)
			if memory[0] == output {
				noun = i
				verb = j
				break Loop
			}
		}
	}
	return noun, verb
}

func main() {
	const filename = "../input.txt"
	const desired = 19690720

	buf, _ := ioutil.ReadFile(filename)
	input := string(buf)

	sli := toIntSlice(input)
	noun, verb := checkForDesiredOutput(sli, desired)
	fmt.Println("100 * noun + verb:", 100*noun+verb)
}
