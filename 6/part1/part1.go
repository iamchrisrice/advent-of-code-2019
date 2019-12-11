package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Object struct {
	orbits string
}

type OrbitMap = map[string]Object

func getStringSliceFromFile(filename string) []string {
	var strings []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func addToOrbitMap(orbitMap OrbitMap, object string, orbits string) {
	if _, ok := orbitMap[orbits]; !ok {
		orbitMap[orbits] = Object{}
	}
	if _, ok := orbitMap[object]; ok && orbitMap[object].orbits != "" {
		panic("object can only have one orbit")
	}
	orbitMap[object] = Object{orbits}
}

func getOrbits(orbitMap OrbitMap, object string) []string {
	orbits := make([]string, 0)
	for {
		object = orbitMap[object].orbits
		if _, ok := orbitMap[object]; !ok {
			break
		}
		orbits = append([]string{object}, orbits...)
	}
	return orbits
}

func countOrbits(orbitMap OrbitMap, object string) int {
	return len(getOrbits(orbitMap, object))
}

func countAllOrbits(orbitMap OrbitMap) int {
	count := 0
	for object := range orbitMap {
		count = count + countOrbits(orbitMap, object)
	}
	return count
}

func main() {
	const filename = "../input.txt"

	input := getStringSliceFromFile(filename)

	orbitMap := make(OrbitMap)

	for _, line := range input {
		objects := strings.Split(line, ")")
		addToOrbitMap(orbitMap, objects[1], objects[0])
	}

	fmt.Println("Total number of orbits in map =", countAllOrbits(orbitMap))

}
