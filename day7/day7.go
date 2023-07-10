package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []*Elm

func (s *stack) push(e *Elm) {
	*s = append(*s, e)
}

func (s *stack) pop() *Elm {
	e := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return e
}

const (
	FILE = "FILE"
	DIR  = "DIR"
)

type Elm struct {
	name     string
	elmType  string
	parent   *Elm
	children []*Elm
	size     int
}

func NewElm(name, elmType string, parent *Elm) *Elm {
	return &Elm{
		name:    name,
		elmType: elmType,
		parent:  parent,
	}
}

func (e *Elm) AddChild(child *Elm) {
	e.children = append(e.children, child)
}

func (e *Elm) HasChildren() int {
	return len(e.children)
}

func (e *Elm) Size() int {
	s := e.size
	for _, c := range e.children {
		s += c.Size()
	}
	return s
}

func part1(sc *bufio.Scanner) {

	root := NewElm("/", DIR, nil)
	p := root

	// skip first line
	sc.Scan()
	for sc.Scan() {
		if strings.Index(sc.Text(), "$ cd") == 0 {
			var path string
			fmt.Sscanf(sc.Text(), "$ cd %s", &path)
			if path == ".." {
				p = p.parent
			} else {
				newElm := NewElm(path, DIR, p)
				p.AddChild(newElm)
				p = newElm
			}
		} else if _, err := strconv.Atoi(string(sc.Text()[0])); err == nil {
			var filename string
			var size int
			fmt.Sscanf(sc.Text(), "%d %s", &size, &filename)
			newElm := NewElm(filename, FILE, p)
			newElm.size = size
			p.AddChild(newElm)
		}
	}

	// calc
	var st stack
	res := 0
	st.push(root)
	for len(st) > 0 {
		p := st.pop()

		size := p.Size()
		if p.elmType == DIR && p.name != "/" && size < 100000 {
			res += size
		}

		for _, child := range p.children {
			st.push(child)
		}
	}

	fmt.Println(res)

}

func part2(sc *bufio.Scanner) {
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if os.Args[1] == "1" {
		part1(sc)
	} else if os.Args[1] == "2" {
		part2(sc)
	}

}
