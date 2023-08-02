package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Queue []int

func NewQueue() Queue {
	return Queue(make([]int, 0))
}

func (q *Queue) Push(a int) {
	*q = append(*q, a)
}

func (q *Queue) Pop() int {
	tmp := (*q)[0]
	*q = (*q)[1:]
	return tmp
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

type Op func(int) int
type Test func(int) bool

type Monkey struct {
	items     Queue
	operation Op
	test      Test

	ml *MonkeysList
	m1 int // index of 1st monkey in the ml
	m2 int // index of 2nd monkey in the ml
}

func (m *Monkey) StartThrowing() int {
	ml := m.ml
	numOfInspection := 0
	for !m.items.IsEmpty() {
		numOfInspection++
		top := m.items.Pop()
		top = m.operation(top)
		top = int(math.Round(float64(top / 3)))
		if m.test(top) {
			ml.GetMonkey(m.m1).items.Push(top)
		} else {
			ml.GetMonkey(m.m2).items.Push(top)
		}
	}

	return numOfInspection
}

// to help with debugging
func (m *Monkey) String() string {
	return fmt.Sprintf("items = %+v", m.items)
}

type MonkeysList []*Monkey

func NewMonkeyList() MonkeysList {
	return MonkeysList(make([]*Monkey, 0))
}

func (ml *MonkeysList) CreateNewMonkey() *Monkey {
	m := &Monkey{}
	m.ml = ml
	*ml = append(*ml, m)
	return m
}

func (ml MonkeysList) GetMonkey(index int) *Monkey {
	return ml[index]
}

func (ml *MonkeysList) StartRounds() {
	numOfInspections := make([]int, len(*ml))
	for round := 0; round < 20; round++ {
		for i, m := range *ml {
			numOfInspections[i] = m.StartThrowing() + numOfInspections[i]
		}
	}

	sort.IntSlice(numOfInspections).Sort()
	top2 := numOfInspections[len(*ml)-2:]
	fmt.Println(top2[0] * top2[1])
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
	items := line[strings.Index(line, ":")+2:]
	parsedItems := make([]int, 0)
	for _, item := range strings.Split(items, ", ") {
		v, _ := strconv.Atoi(item)
		parsedItems = append(parsedItems, v)
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

func ParseMonkeys(sc *bufio.Scanner) MonkeysList {
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
		} else if strings.Index(line, "If false") == 0 {
			var i int
			fmt.Sscanf(line, "If false: throw to monkey %d\n", &i)
			currentMonkey.m2 = i
		}

	}

	return ml
}

func part1(sc *bufio.Scanner) {
	ml := ParseMonkeys(sc)
	ml.StartRounds()
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
