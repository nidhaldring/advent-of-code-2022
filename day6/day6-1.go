package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func uniq(s string) bool {
	for i, ch := range s {
		if strings.Index(s, string(ch)) != i {
			return false
		}
	}
	return true
}

func part1(sc *bufio.Scanner) {
	sc.Scan()
	line := sc.Text()
	for i := 3; i < len(line)-1; i++ {
		if uniq(line[i-3 : i+1]) {
			fmt.Println(i + 1)
			return
		}
	}
}

func part2(sc *bufio.Scanner) {
	sc.Scan()
	line := sc.Text()
	for i := 13; i < len(line)-1; i++ {
		if uniq(line[i-13 : i+1]) {
			fmt.Println(i + 1)
			return
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	if os.Args[1] == "1" {
		part1(sc)
	} else if os.Args[1] == "2" {
		part2(sc)
	}
}
