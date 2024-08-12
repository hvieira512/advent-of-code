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

func getFileContent(input string) (string, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func getPart1(input string) (result int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result += calculateFuel(mass)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func getPart2(input string) (result int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel := 0
		for currentMass := mass; currentMass > 0; currentMass = calculateFuel(currentMass) {
			currentFuel := calculateFuel(currentMass)
			if currentFuel > 0 {
				fuel += currentFuel
			}
		}
		result += fuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	fmt.Println("---- Day 1 ----")

	data, err := getFileContent("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	resultPart1 := getPart1(data)
	fmt.Printf("Result Part 1: %d\n", resultPart1)

	resultPart2 := getPart2(data)
	fmt.Printf("Result Part 2: %d", resultPart2)
}
