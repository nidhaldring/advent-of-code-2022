package main

import "fmt"

type Queue []*PathPoint

func NewQueue() *Queue {
	return &Queue{}
}

func (s *Queue) Push(c *PathPoint) {
	*s = append(*s, c)
}

func (s *Queue) Pop() *PathPoint {
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}

func (s *Queue) Reverse() *Queue {
	rev := NewQueue()
	for i := len(*s) - 1; i >= 0; i-- {
		rev.Push((*s)[i])
	}
	return rev
}

func (s *Queue) Contains(n *PathPoint) bool {
	for _, v := range *s {
		if v.x == n.x && v.y == n.y {
			return true
		}
	}
	return false
}

func (s *Queue) Empty() bool {
	return len(*s) == 0
}

func (s *Queue) Len() int {
	return len(*s)
}

func (s *Queue) PrintPath() {
	for v, p := range *s {
		fmt.Printf("(%d, %d) ", p.x, p.y)
		if v != len(*s)-1 {
			fmt.Print("-> ")
		}
	}
	fmt.Printf(" Path Len = %d\n\n", len(*s)-1)
}
