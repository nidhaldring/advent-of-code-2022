package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1(sc *bufio.Scanner) {
	cyclesNb := 0
	x := 1
	sumOfSignals := 0
	for sc.Scan() {
		var end int
		if strings.Index(sc.Text(), "noop") == 0 {
			end = 1
		} else {
			end = 2
		}

		for i := 0; i < end; i++ {
			cyclesNb++

			if cyclesNb%40 == 20 {
				sumOfSignals += x * cyclesNb
			}

			if i == 1 {
				var toAdd int
				fmt.Sscanf(sc.Text(), "addx %d\n", &toAdd)
				x += toAdd
			}
		}

	}

	fmt.Println(sumOfSignals)

}

func part2(sc *bufio.Scanner) {
	m := make([]string, 40*6)
	cyclesNb := 0
	x := 1
	for sc.Scan() {
		var end int
		if strings.Index(sc.Text(), "noop") == 0 {
			end = 1
		} else {
			end = 2
		}

		for i := 0; i < end; i++ {
			if cyclesNb%40 >= x-1 && cyclesNb%40 <= x+1 {
				m[cyclesNb] = "#"
			} else {
				m[cyclesNb] = "."
			}

			if i == 1 {
				var toAdd int
				fmt.Sscanf(sc.Text(), "addx %d\n", &toAdd)
				x += toAdd
			}
			cyclesNb++
		}

	}

	for i := 0; i < 40*6; i++ {
		if i%40 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%s", m[i])
	}
	fmt.Printf("\n")

}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	if os.Args[1] == "1" {
		part1(sc)
	} else if os.Args[1] == "2" {
		part2(sc)
	}
}
