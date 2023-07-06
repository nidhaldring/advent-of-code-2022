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
		switch line[2] {
		case 'Y':
			score += int(line[0]-'A') + 1
			score += 3
		case 'X':
			switch line[0] {
			case 'A':
				score += 3
			case 'B':
				score += 1
			case 'C':
				score += 2
			}
		case 'Z':
			score += 6
			switch line[0] {
			case 'A':
				score += 2
			case 'B':
				score += 3
			case 'C':
				score += 1
			}
		}
	}

	fmt.Println(score)
}
