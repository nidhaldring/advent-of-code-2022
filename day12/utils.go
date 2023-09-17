package main

import (
	"bufio"
	"cmp"
	"math"
	"slices"
	"strings"
)

func GetNeighbors(n, goal *PathPoint, matrix [][]string) []*PathPoint {
	maxX := len(matrix)
	maxY := len(matrix[0])

	res := make([]*PathPoint, 0)
	if n.x > 0 && minus(matrix[n.x-1][n.y], matrix[n.x][n.y]) <= 1 {
		res = append(res, &PathPoint{x: n.x - 1, y: n.y, parent: n})
	}

	if n.x < maxX-1 && minus(matrix[n.x+1][n.y], matrix[n.x][n.y]) <= 1 {
		res = append(res, &PathPoint{x: n.x + 1, y: n.y, parent: n})
	}

	if n.y > 0 && minus(matrix[n.x][n.y-1], matrix[n.x][n.y]) <= 1 {
		res = append(res, &PathPoint{x: n.x, y: n.y - 1, parent: n})
	}

	if n.y < maxY-1 && minus(matrix[n.x][n.y+1], matrix[n.x][n.y]) <= 1 {
		res = append(res, &PathPoint{x: n.x, y: n.y + 1, parent: n})
	}

	// sort neighbors by closest to goal
	slices.SortFunc(res, func(a, b *PathPoint) int {
		aDistToGoal := math.Sqrt(math.Pow(float64(a.x)-float64(goal.x), 2) + math.Pow(float64(a.y)-float64(goal.y), 2))
		bDistToGoal := math.Sqrt(math.Pow(float64(b.x)-float64(goal.x), 2) + math.Pow(float64(b.y)-float64(goal.y), 2))
		return cmp.Compare(bDistToGoal, aDistToGoal)
	})

	return res
}

func ParsePart1Input(sc *bufio.Scanner) ([][]string, *PathPoint, *PathPoint) {
	var start *PathPoint
	var goal *PathPoint
	var matrix [][]string
	for line := 0; sc.Scan(); line++ {
		startIndex := strings.Index(sc.Text(), "S")
		if startIndex != -1 {
			start = &PathPoint{x: line, y: startIndex}
		}

		goalIndex := strings.Index(sc.Text(), "E")
		if goalIndex != -1 {
			goal = &PathPoint{x: line, y: goalIndex}
		}
		matrix = append(matrix, strings.Split(sc.Text(), ""))
	}

	return matrix, start, goal
}

func ParsePart2Input(sc *bufio.Scanner) ([][]string, []*PathPoint, *PathPoint) {
	startPoints := make([]*PathPoint, 0)
	var goal *PathPoint
	var matrix [][]string
	for line := 0; sc.Scan(); line++ {
		startIndex := strings.Index(sc.Text(), "S")
		if startIndex != -1 {
			startPoints = append(startPoints, &PathPoint{x: line, y: startIndex})
		}

		startIndex = strings.Index(sc.Text(), "a")
		if startIndex != -1 {
			startPoints = append(startPoints, &PathPoint{x: line, y: startIndex})
		}

		goalIndex := strings.Index(sc.Text(), "E")
		if goalIndex != -1 {
			goal = &PathPoint{x: line, y: goalIndex}
		}
		matrix = append(matrix, strings.Split(sc.Text(), ""))
	}

	return matrix, startPoints, goal
}

func minus(a, b string) int {
	if a == "S" {
		a = "a"
	}

	if b == "S" {
		b = "a"
	}

	if a == "E" {
		a = "z"
	}

	if b == "E" {
		b = "z"
	}

	return int([]byte(a)[0]) - int([]byte(b)[0])
}
