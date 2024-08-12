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

func getPart1(input string) (result int) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel := int(math.Floor(float64(mass/3))) - 2
		result += fuel
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

		fmt.Printf("Mass:%d | Fuel:", mass)

		totalFuel := 0
		for {
			fuel := int(math.Floor(float64(mass/3))) - 2
			fmt.Printf("%d ", fuel)
			if (int(math.Floor(float64(fuel/3))) - 2) <= 0 {
				if totalFuel == 0 {
					totalFuel = fuel
				} else {
					totalFuel += fuel
				}
				break
			}
			mass = fuel
			totalFuel += fuel
		}
		fmt.Printf("= %d\n", totalFuel)
		result += totalFuel
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
