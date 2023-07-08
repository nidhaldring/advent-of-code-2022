package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	res := 0

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		for _, ch := range firstHalf {
			if strings.Contains(secondHalf, string(ch)) {
				if ch >= 'a' {
					res += int(ch-'a') + 1
				} else {
					res += int(ch-'A') + 27
				}
				break
			}
		}
	}

	fmt.Println(res)
}
