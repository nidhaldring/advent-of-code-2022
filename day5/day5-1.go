package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stackElm struct {
	value string
	next  *stackElm
}

type stack struct {
	head *stackElm
	tail *stackElm
}

func newStack() *stack {
	return &stack{
		head: nil,
		tail: nil,
	}
}

func (s *stack) push(value string) {
	elm := &stackElm{value: value, next: nil}
	if s.head == nil && s.tail == nil {
		s.tail = elm
	} else {
		elm.next = s.head
	}

	s.head = elm
}

func (s *stack) pop() string {
	if s.head == s.tail {
		s.tail = nil
	}

	tmp := s.head
	s.head = tmp.next
	tmp.next = nil

	return tmp.value
}

func (s *stack) reverse() {
	var o, p, q *stackElm = nil, s.head, s.head.next
	for q != nil {
		p.next = o
		o = p
		p = q
		q = q.next
	}
	p.next = o

	tmp := s.head
	s.head = s.tail
	s.tail = tmp
}

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

		for quantity > 0 {
			stacks[dest-1].push(stacks[index-1].pop())
			quantity--
		}

	}

	for _, stack := range stacks {
		fmt.Print(stack.pop())
	}
	fmt.Print("\n")

}
