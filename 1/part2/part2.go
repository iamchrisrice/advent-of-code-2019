package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func fuelRequirement(mass float64) float64 {
	return math.Floor(mass/3) - 2
}

func additionalFuelRequirement(fuel float64) float64 {
	var additionalFuel float64
	f := fuel
	for {
		f = fuelRequirement(f)
		if f <= 0 {
			break
		}
		additionalFuel = additionalFuel + f
	}
	return additionalFuel
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
		moduleFuel := fuelRequirement(moduleMass)
		additionalFuel := additionalFuelRequirement(moduleFuel)
		spacecraftFuelRequirement = spacecraftFuelRequirement + moduleFuel + additionalFuel
	}

	fmt.Println(int(spacecraftFuelRequirement))
}
