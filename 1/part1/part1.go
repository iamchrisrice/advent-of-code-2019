package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func moduleFuelRequirement(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func main() {
	const filename = "../input.txt"

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var spacecraftFuelRequirement float64
	for scanner.Scan() {
		moduleMass, _ := strconv.ParseFloat(scanner.Text(), 64)
		moduleFuel := moduleFuelRequirement(moduleMass)
		spacecraftFuelRequirement = spacecraftFuelRequirement + moduleFuel
	}

	fmt.Println(int(spacecraftFuelRequirement))
}
