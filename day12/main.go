package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func findShortestPath(matrix [][]string, start, goal *PathPoint) int {
	// start traversing the matrix searching for the least num of steps to reach E
	nodes := NewQueue()
	visitedNodes := make(map[[2]int]bool)

	nodes.Push(start)
	for !nodes.Empty() {
		node := nodes.Pop()
		neighbors := GetNeighbors(node, goal, matrix)
		for _, n := range neighbors {
			if n.x == goal.x && n.y == goal.y {
				// -1 to not include the goal
				return n.ConstructPath().Len() - 1
			} else if !visitedNodes[[2]int{n.x, n.y}] {
				visitedNodes[[2]int{n.x, n.y}] = true
				nodes.Push(n)
			}
		}
	}

	return 0
}

func part1(sc *bufio.Scanner) {
	matrix, start, goal := ParsePart1Input(sc)
	fmt.Printf("The shortest path is %d\n", findShortestPath(matrix, start, goal))
}

func part2(sc *bufio.Scanner) {
	matrix, startPoints, goal := ParsePart2Input(sc)

	min := findShortestPath(matrix, startPoints[0], goal)
	for i := 1; i < len(startPoints); i++ {
		steps := findShortestPath(matrix, startPoints[i], goal)
		if steps < min {
			min = steps
		}
	}

	fmt.Println(min)
}

func main() {
	part := flag.Int("part", 1, "1 or 2 to specify part")
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)
	if *part == 1 {
		part1(sc)
	} else {
		part2(sc)
	}
}
