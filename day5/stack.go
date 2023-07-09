package main

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

func (s *stack) String() string {
	res := "[ "

	p := s.head
	for p != nil {
		res += p.value + " "
		p = p.next
	}

	res += "]"
	return res
}
