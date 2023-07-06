package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	score := 0
	for sc.Scan() {
		line := sc.Text()
		score += int(line[2]-'X') + 1
		switch line {
		case "A Y",
			"B Z",
			"C X":
			score += 6
		case "A X",
			"B Y",
			"C Z":
			score += 3
		}
	}

	fmt.Println(score)
}
