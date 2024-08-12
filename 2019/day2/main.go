package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2019/utils"
)

func getPart1(data string) (result int) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		row, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(row)

		// Do logic here
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func getPart2(data string) (result int) {
	scanner := bufio.NewScanner(strings.NewReader(data))

	for scanner.Scan() {
		row, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(row)

		// Do logic here
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	fmt.Println("---- Day 1 ----")

	data, err := utils.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	resultPart1 := getPart1(data)
	fmt.Printf("Result Part 1: %d\n", resultPart1)

	resultPart2 := getPart2(data)
	fmt.Printf("Result Part 2: %d", resultPart2)

}
