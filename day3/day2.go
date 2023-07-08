package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func common(str1, str2 string) string {
	res := ""
	for _, ch := range str1 {
		if strings.Contains(str2, string(ch)) && !strings.Contains(res, string(ch)) {
			res += string(ch)
		}
	}
	return res
}

// This only works assuming the input is correct
func main() {

	res := 0
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line1 := sc.Text()

		sc.Scan()
		line2 := sc.Text()

		sc.Scan()
		line3 := sc.Text()

		ch := common(line1, common(line2, line3))[0]

		if ch >= 'a' {
			res += int(ch-'a') + 1
		} else {
			res += int(ch-'A') + 27
		}
	}

	fmt.Println(res)

}
