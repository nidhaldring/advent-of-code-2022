package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(sc *bufio.Scanner) {
	matrix, start, goal := ParseInput(sc)

	// start traversing the matrix searching for the least num of steps to reach E
	nodes := NewStack()
	visitedNodes := NewStack()
	nodes.Push(start)
LOOP:
	for !nodes.Empty() {
		node := nodes.Pop()
		visitedNodes.Push(node)

		neighbors := GetNeighbors(node, goal, matrix)
		for _, n := range neighbors {
			if n.x == goal.x && n.y == goal.y {
				n.ConstructPath().Reverse().PrintPath()
				break LOOP
			} else if !visitedNodes.Contains(n) {
				nodes.Push(n)
			}
		}
	}
}

func part2(sc *bufio.Scanner) {}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify the part, either 1 or 2!")
		return
	}

	sc := bufio.NewScanner(os.Stdin)
	if os.Args[1] == "1" {
		part1(sc)
	} else {
		part2(sc)
	}
}
