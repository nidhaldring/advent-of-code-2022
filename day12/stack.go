package main

import "fmt"

type Stack []*PathPoint

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(c *PathPoint) {
	*s = append(*s, c)
}

func (s *Stack) Pop() *PathPoint {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
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
