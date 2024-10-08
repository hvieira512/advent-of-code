package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func GetLinesFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	return lines
}

func GetValuesLine(lines []string, lineNumber int, sep string) ([]int, error) {
	values := []int{}
	valuesAux := lines[lineNumber]

	valuesAux = strings.ReplaceAll(valuesAux, sep, " ")
	valuesAux = strings.Trim(valuesAux, " ")
	valuesStr := strings.Fields(valuesAux)

	for _, timeStr := range valuesStr {
		aux, err := strconv.Atoi(timeStr)
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, aux)
	}

	return values, nil
}
