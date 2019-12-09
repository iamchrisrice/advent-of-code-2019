package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Wire struct {
	id           int
	instructions string
}

type Wires = map[Wire]bool

type Grid = map[Position]Wires

func instructionSlice(instructions string) []string {
	return strings.Split(instructions, ",")
}

func wire(wire Wire, grid Grid) {
	var pos = Position{0, 0}
	for _, instruction := range instructionSlice(wire.instructions) {
		direction := instruction[0:1]
		distance, _ := strconv.Atoi(instruction[1:])
		switch direction {
		case "U":
			var next Position
			for i := 0; i < distance; i++ {
				next = Position{pos.x, pos.y + 1}
				if len(grid[next]) == 0 {
					grid[next] = make(Wires)
				}
				grid[next][wire] = true
				pos = next
			}
		case "D":
			var next Position
			for i := 0; i < distance; i++ {
				next = Position{pos.x, pos.y - 1}
				if len(grid[next]) == 0 {
					grid[next] = make(Wires)
				}
				grid[next][wire] = true
				pos = next
			}
		case "R":
			var next Position
			for i := 0; i < distance; i++ {
				next = Position{pos.x + 1, pos.y}
				if len(grid[next]) == 0 {
					grid[next] = make(Wires)
				}
				grid[next][wire] = true
				pos = next
			}
		case "L":
			var next Position
			for i := 0; i < distance; i++ {
				next = Position{pos.x - 1, pos.y}
				if len(grid[next]) == 0 {
					grid[next] = make(Wires)
				}
				grid[next][wire] = true
				pos = next
			}
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

func getDistance(position Position) int {
	return int(math.Abs(float64(position.x)) + math.Abs(float64(position.y)))
}

func getIntersectionDistances(grid Grid) []int {
	var distances []int
	for _, pos := range findIntersections(grid) {
		distances = append(distances, getDistance(pos))
	}
	return distances
}

func getShortestIntersectionDistance(grid Grid) int {
	var shortest int
	for index, distance := range getIntersectionDistances(grid) {
		if index == 0 || distance < shortest {
			shortest = distance
		}
	}
	return shortest
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
		w := Wire{id, scanner.Text()}
		wire(w, grid)
	}

	fmt.Println("Manhattan distance:", getShortestIntersectionDistance(grid))
}