package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Wire struct {
	id int
}

type Wires = map[Wire][]int

type Grid = map[Position]Wires

func instructionSlice(instructions string) []string {
	return strings.Split(instructions, ",")
}

func wire(wire Wire, instructions string, grid Grid) {
	var pos = Position{0, 0}
	var steps = 0
	for _, instruction := range instructionSlice(instructions) {
		direction := instruction[0:1]
		distance, _ := strconv.Atoi(instruction[1:])
		var next Position
		for i := 0; i < distance; i++ {
			steps++
			switch direction {
			case "U":
				next = Position{pos.x, pos.y + 1}
			case "D":
				next = Position{pos.x, pos.y - 1}
			case "R":
				next = Position{pos.x + 1, pos.y}
			case "L":
				next = Position{pos.x - 1, pos.y}
			}
			if len(grid[next]) == 0 {
				grid[next] = make(Wires)
			}
			grid[next][wire] = append(grid[next][wire], steps)
			pos = next
		}
	}
}

func findIntersections(grid Grid) []Position {
	var intersections []Position

	for pos, wires := range grid {
		if len(wires) > 1 {
			intersections = append(intersections, pos)
		}
	}
	return intersections
}

func lowestInt(sli []int) int {
	var shortest int
	for _, current := range sli {
		if shortest == 0 || current < shortest {
			shortest = current
		}
	}
	return shortest
}

func getShortestTotalStepsToPosition(grid Grid, position Position) int {
	var steps int
	for _, y := range grid[position] {
		steps = steps + lowestInt(y)
	}
	return steps
}

func getShortestStepsToAnyIntersection(grid Grid) int {
	var stepsToIntersections []int
	for _, y := range findIntersections(grid) {
		stepsToIntersections = append(stepsToIntersections, getShortestTotalStepsToPosition(grid, y))
	}
	return lowestInt(stepsToIntersections)
}

func main() {
	const filename = "../input.txt"

	var grid = make(Grid)
	var id = 0

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		id++
		w := Wire{id}
		wire(w, scanner.Text(), grid)
	}

	fmt.Println("Least steps:", getShortestStepsToAnyIntersection(grid))

}
