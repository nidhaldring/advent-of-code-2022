package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func parseInput() [][]string {
	res := make([][]string, 0)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		res = append(res, strings.Split(sc.Text(), ""))
	}

	return res
}

func part1() {
	matrix := parseInput()

	visible := len(matrix)*2 + len(matrix[0])*2 - 4
	for i := 1; i < len(matrix)-1; i++ {
	LOOP:
		for j := 1; j < len(matrix[0])-1; j++ {
			elm := matrix[i][j]

			var k int
			for k = i - 1; k >= 0; k-- {
				if matrix[k][j] >= elm {
					break
				}
			}
			if k < 0 {
				visible++
				continue LOOP
			}

			for k = i + 1; k < len(matrix); k++ {
				if matrix[k][j] >= elm {
					break
				}
			}
			if k == len(matrix) {
				visible++
				continue LOOP
			}

			for k = j - 1; k >= 0; k-- {
				if matrix[i][k] >= elm {
					break
				}
			}
			if k < 0 {
				visible++
				continue LOOP
			}

			for k = j + 1; k < len(matrix[0]); k++ {
				if matrix[i][k] >= elm {
					break
				}
			}
			if k == len(matrix[0]) {
				visible++
				continue LOOP
			}
		}
	}

	fmt.Println(visible)
}

func main() {
	part := flag.Int("part", 1, "choose either 1 or 2")
	flag.Parse()

	if *part == 1 {
		part1()
	} else {

	}
}
