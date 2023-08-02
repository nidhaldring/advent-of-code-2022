package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

type Op func(int) int
type Test func(int) bool

type Monkey struct {
	items     Stack
	operation Op
	test      Test

	ml *MonkeysList
	m1 int // index of 1st monkey in the ml
	m2 int // index of 2nd monkey in the ml
}

func (m *Monkey) StartThrowing() {
	ml := m.ml
	for !m.items.IsEmpty() {
		top := m.operation(m.items.Pop())
		top = int(math.Round(float64(top) / 3))
		if m.test(top) {
			ml.GetMonkey(m.m1).items.Push(top)
		} else {
			ml.GetMonkey(m.m2).items.Push(top)
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

func (ml *MonkeysList) GetMonkey(index int) *Monkey {
	return ml.monkeys[index-1]
}

func ApplyOperand(op string, x, y int) int {
	switch op {
	case "+":
		return x + y
	case "*":
		return x * y
	case "-":
		return x - y
	default:
		return x / y
	}
}

func ParseStartingItems(line string) []int {
	var items string
	fmt.Sscanf(line, "Starting Items: %s\n", &items)

	parsedItems := NewStack()
	for _, item := range strings.Split(items, ",") {
		v, _ := strconv.Atoi(item)
		parsedItems.Push(v)
	}
	return parsedItems
}

func ParseOperation(line string) Op {
	var right, op string
	fmt.Sscanf(line, "Operation: new = old %s %s\n", &op, &right)

	return func(i int) int {
		if right == "old" {
			return ApplyOperand(op, i, i)
		}

		v, _ := strconv.Atoi(right)
		return ApplyOperand(op, i, v)
	}
}

func ParseTest(line string) Test {
	var div int
	fmt.Sscanf(line, "Test: divisible by %d", &div)

	return func(i int) bool {
		return i%div == 0
	}

}

func ParseMonkeys(sc *bufio.Scanner) *MonkeysList {
	ml := NewMonkeyList()
	var currentMonkey *Monkey
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.Index(line, "Monkey") == 0 {
			currentMonkey = ml.CreateNewMonkey()
		} else if strings.Index(line, "Starting") == 0 {
			currentMonkey.items = ParseStartingItems(line)
		} else if strings.Index(line, "Operation") == 0 {
			currentMonkey.operation = ParseOperation(line)
		} else if strings.Index(line, "Test") == 0 {
			currentMonkey.test = ParseTest(line)
		} else if strings.Index(line, "If true") == 0 {
			var i int
			fmt.Sscanf(line, "If true: throw to monkey %d\n", &i)
			currentMonkey.m1 = i
		} else if strings.Index(line, "Test") == 0 {
			var i int
			fmt.Sscanf(line, "If false: throw to monkey %d\n", &i)
			currentMonkey.m1 = i
		}

	}

	return ml
}

func part1(sc *bufio.Scanner) {
	ml := ParseMonkeys(sc)
	fmt.Println(ml)
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
