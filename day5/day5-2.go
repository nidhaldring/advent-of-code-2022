package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	stacks := make([]*stack, 0)
	sc.Scan()
	for i := 0; i < len(sc.Text()); i += 4 {
		stacks = append(stacks, newStack())
	}

	for !strings.Contains(sc.Text(), " 1") {
		line := sc.Text()
		for i, j := 0, 0; i <= len(line)-3; i, j = i+4, j+1 {
			if line[i:i+3] != "   " {
				stacks[j].push(line[i+1 : i+2])
			}
		}

		sc.Scan()
	}

	sc.Scan() // skip empty line

	// reverse stacks order
	for i := 0; i < len(stacks); i++ {
		stacks[i].reverse()
	}

	// read instructions
	for sc.Scan() {
		var quantity, index, dest int
		fmt.Sscanf(sc.Text(), "move %d from %d to %d", &quantity, &index, &dest)

		v := ""
		for i := quantity - 1; i >= 0; i-- {
			v += stacks[index-1].pop()
		}

		for i := len(v) - 1; i >= 0; i-- {
			stacks[dest-1].push(string(v[i]))
		}
	}

	for _, s := range stacks {
		fmt.Println(s)
	}

	for _, stack := range stacks {
		fmt.Print(stack.pop())
	}
	fmt.Print("\n")

}
