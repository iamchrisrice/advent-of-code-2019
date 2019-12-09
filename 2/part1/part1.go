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

func main() {
	const filename = "../input.txt"

	buf, _ := ioutil.ReadFile(filename)
	input := string(buf)

	sli := toIntSlice(input)
	sli[1] = 12
	sli[2] = 2
	emulate(sli)
	fmt.Println("Position 0:", sli[0])
}
