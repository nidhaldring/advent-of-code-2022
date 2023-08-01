package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Op func(int) int
type Test func(int) bool

type Stack []int

func NewStack() Stack {
	return Stack(make([]int, 0))
}

func (s *Stack) Push(a int) {
	*s = append(*s, a)
}

func (s *Stack) Pop() int {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

type Monkey struct {
	items     Stack
	operation Op
	test      Test

	ml *MonkeysList
	m1 int
	m2 int
}

func (m *Monkey) SetItems(items string) {
	parsedItems := NewStack()
	for _, v := range strings.Split(items, ",") {
		i, _ := strconv.Atoi(v)
		parsedItems.Push(i)
	}
	m.items = parsedItems
}

func (m *Monkey) StartThrowing() {
	for !m.items.IsEmpty() {
		top := m.items.Pop()
		top = m.operation(top)
		top = int(math.Round(float64(top) / 3))
		if m.test(top) {
			m.ml.Get(m.m1).items.Push(top)
		} else {
			m.ml.Get(m.m2).items.Push(top)
		}
	}
}

type MonkeysList struct {
	monkeys []*Monkey
}

func NewMonkeyList() *MonkeysList {
	return &MonkeysList{
		monkeys: make([]*Monkey, 0),
	}
}

func (ml *MonkeysList) CreateNewMonkey() *Monkey {
	m := &Monkey{}
	m.ml = ml
	ml.monkeys = append(ml.monkeys, m)
	return m
}

func (ml *MonkeysList) Get(index int) *Monkey {
	return ml.monkeys[index-1]
}

func part1(sc *bufio.Scanner) {
	ml := NewMonkeyList()
	var currentMonkey *Monkey
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.Index(line, "Monkey") == 0 {
			currentMonkey = ml.CreateNewMonkey()
		} else if strings.Index(line, "Starting") == 0 {
			var items string
			fmt.Sscanf(line, "Starting Items: %s\n", &items)
			currentMonkey.SetItems(items)
		} else if strings.Index(line, "Operation") == 0 {
			var items string
			fmt.Sscanf(line, "Starting Items: %s\n", &items)
			currentMonkey.SetItems(items)
		}

	}
}

func part2(sc *bufio.Scanner) {}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if os.Args[1] == "1" {
		part1(sc)
	} else if os.Args[1] == "2" {
		part2(sc)
	}
}
