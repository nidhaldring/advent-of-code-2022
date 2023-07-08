package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	res := 0
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		var fistRangeFirstBound, fistRangeSecondBound, secondRangeFirstBound, secondRangeSecondBound int
		fmt.Sscanf(sc.Text(),
			"%d-%d,%d-%d",
			&fistRangeFirstBound, &fistRangeSecondBound, &secondRangeFirstBound, &secondRangeSecondBound)

		if (fistRangeFirstBound <= secondRangeFirstBound && fistRangeSecondBound >= secondRangeSecondBound) ||
			(secondRangeFirstBound <= fistRangeFirstBound && secondRangeSecondBound >= fistRangeSecondBound) ||
			(fistRangeSecondBound >= secondRangeFirstBound && fistRangeFirstBound <= secondRangeSecondBound) ||
			(secondRangeFirstBound <= fistRangeSecondBound && secondRangeFirstBound >= fistRangeFirstBound) {
			res += 1
		}
	}

	fmt.Println(res)
}
