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

func part1() {
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

func part2() {
	arr := [3]int{0, 0, 0}
	for _, elfCal := range strings.Split(readInput(), "\n\n") {
		sum := 0
		for _, cal := range strings.Split(elfCal, "\n") {
			n, _ := strconv.Atoi(cal)
			sum += n
		}

		if sum > arr[0] {
			arr[0], arr[1], arr[2] = sum, arr[0], arr[1]
		} else if sum > arr[1] {
			arr[1], arr[2] = sum, arr[1]
		} else if sum > arr[2] {
			arr[2] = sum
		}
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	fmt.Println(sum)

}

func main() {
	part1()
	part2()
}
