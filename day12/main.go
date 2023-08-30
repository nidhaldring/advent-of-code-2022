package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type PathPoint struct {
	x, y   int
	parent *PathPoint
}

type Stack []*PathPoint

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(c *PathPoint) {
	*s = append(*s, c)
}

func (s *Stack) Pop() *PathPoint {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}

func (s *Stack) Reverse() *Stack {
	rev := NewStack()
	for i := len(*s) - 1; i >= 0; i-- {
		rev.Push((*s)[i])
	}
	return rev
}

func (s *Stack) Contains(n *PathPoint) bool {
	for _, v := range *s {
		if v.x == n.x && v.y == n.y {
			return true
		}
	}
	return false
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) PrintPath() {
	for v, p := range *s {
		fmt.Printf("(%d, %d) ", p.x, p.y)
		if v != len(*s)-1 {
			fmt.Print("-> ")
		}
	}
	fmt.Printf(" Path Len = %d\n\n", len(*s))
}

func Minus(a, b string) int {
	return int([]byte(a)[0]) - int([]byte(b)[0])
}

func GetNeighbors(n, goal *PathPoint, matrix [][]string) []*PathPoint {
	maxX := len(matrix)
	maxY := len(matrix[0])

	res := make([]*PathPoint, 0)
	if n.x > 0 && (matrix[n.x-1][n.y] == "E" || Minus(matrix[n.x-1][n.y], matrix[n.x][n.y]) <= 1) {
		res = append(res, &PathPoint{x: n.x - 1, y: n.y, parent: n})
	}

	if n.x < maxX-1 && (matrix[n.x+1][n.y] == "E" || Minus(matrix[n.x+1][n.y], matrix[n.x][n.y]) <= 1) {
		res = append(res, &PathPoint{x: n.x + 1, y: n.y, parent: n})
	}

	if n.y > 0 && (matrix[n.x][n.y-1] == "E" || Minus(matrix[n.x][n.y-1], matrix[n.x][n.y]) <= 1) {
		res = append(res, &PathPoint{x: n.x, y: n.y - 1, parent: n})
	}

	if n.y < maxY-1 && (matrix[n.x][n.y+1] == "E" || Minus(matrix[n.x][n.y+1], matrix[n.x][n.y]) <= 1) {
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

func ConstructPath(p *PathPoint) *Stack {
	path := NewStack()
	for p != nil {
		path.Push(p)
		p = p.parent
	}
	return path
}

func part1(sc *bufio.Scanner) {
	var goal *PathPoint
	var matrix [][]string
	for line := 0; sc.Scan(); line++ {
		goalIndex := strings.Index(sc.Text(), "E")
		if goalIndex != -1 {
			goal = &PathPoint{x: line, y: goalIndex}
		}
		matrix = append(matrix, strings.Split(sc.Text(), ""))
		line++
	}

	// traverse the matrix searching for the least num of steps to reach E
	nodes := NewStack()
	visitedNodes := NewStack()

	nodes.Push(&PathPoint{x: 0, y: 0, parent: nil})

LOOP:
	for !nodes.Empty() {
		node := nodes.Pop()
		visitedNodes.Push(node)

		neighbors := GetNeighbors(node, goal, matrix)
		for _, n := range neighbors {
			if n.x == goal.x && n.y == goal.y {
				ConstructPath(n).Reverse().PrintPath()
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
