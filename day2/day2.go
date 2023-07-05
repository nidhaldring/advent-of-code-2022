package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput() []string {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}

	return strings.Split(strings.TrimSpace(string(f)), "\n")
}

func getRoundScore(line string) int {
	initScore := int(line[2]-'X') + 1

	// A, X => ROCK
	// B, Y => PAPER
	// C, Z => Scissors
	switch line {
	case "A Y",
		"B Z",
		"C X":
		return initScore + 6
	case "A X",
		"B Y",
		"C Z":
		return initScore + 3
	}

	return initScore

}

func main() {
	score := 0
	for _, line := range readInput() {
		score += getRoundScore(line)
	}

	fmt.Println(score)
}
