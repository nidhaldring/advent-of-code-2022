package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	data, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func main() {
	maxCal := 0
	for _, elfCal := range strings.Split(readInput(), "\n\n") {
		sum := 0
		for _, cal := range strings.Split(elfCal, "\n") {
			n, _ := strconv.Atoi(cal)
			sum += n
		}

		if sum > maxCal {
			maxCal = sum
		}
	}

	fmt.Println(maxCal)
}
