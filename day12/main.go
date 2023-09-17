package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1(sc *bufio.Scanner) {
	matrix, start, goal := ParseInput(sc)

	// start traversing the matrix searching for the least num of steps to reach E
	nodes := NewQueue()
	visitedNodes := make(map[[2]int]bool)

	nodes.Push(start)
	for !nodes.Empty() {
		node := nodes.Pop()

		neighbors := GetNeighbors(node, goal, matrix)
		for _, n := range neighbors {
			if n.x == goal.x && n.y == goal.y {
				n.ConstructPath().Reverse().PrintPath()
				break
			} else if !visitedNodes[[2]int{n.x, n.y}] {
				visitedNodes[[2]int{n.x, n.y}] = true
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
