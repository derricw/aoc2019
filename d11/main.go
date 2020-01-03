package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	BLACK   int64                = 0
	WHITE   int64                = 1
	TURNMAP map[string][2]string = map[string][2]string{
		"U": [2]string{"L", "R"},
		"D": [2]string{"R", "L"},
		"L": [2]string{"D", "U"},
		"R": [2]string{"U", "D"},
	}
)

type Surface map[Vec2D]int64

func (s Surface) Draw() {
	var xmin, xmax, ymin, ymax int
	init := false
	for k, _ := range s {
		if !init {
			xmin, xmax, ymin, ymax = k[0], k[0], k[1], k[1]
			init = true
		}
		if xmin > k[0] {
			xmin = k[0]
		}
		if xmax < k[0] {
			xmax = k[0]
		}
		if ymin > k[1] {
			ymin = k[1]
		}
		if ymax < k[1] {
			ymax = k[1]
		}
	}
	for y := ymax; y >= ymin; y-- {
		for x := xmin; x <= xmax; x++ {
			color, _ := s[Vec2D{x, y}]
			if color == 1 {
				fmt.Printf("▓")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

type Robot struct {
	Pos       Vec2D
	Direction string
	Computer  *Process
	Surface   Surface
	Done      bool
}

func NewRobot(program []int64) *Robot {
	return &Robot{
		Pos:       Vec2D{0, 0},
		Direction: "U",
		Computer:  NewProcess(program),
		Surface:   make(Surface),
	}
}

func (r *Robot) Start() {
	r.Done = false
	go func() {
		r.Computer.Start()
		r.Done = true
	}()
}

func (r *Robot) Paint() {
	if r.Surface[r.Pos] == BLACK {
		r.Computer.Input <- BLACK
	} else {
		r.Computer.Input <- WHITE
	}
	color := <-r.Computer.Output
	turn := <-r.Computer.Output

	r.Surface[r.Pos] = color
	r.Turn(turn)
	r.Step()
}

func (r *Robot) Turn(direction int64) {
	r.Direction = TURNMAP[r.Direction][direction]
}

func (r *Robot) Step() {
	r.Pos = r.Pos.Trans(D[r.Direction])
}

func solveP1(in []int64) {
	robot := NewRobot(in)

	robot.Start()
	for !robot.Done {
		robot.Paint()
	}
	output := len(robot.Surface)
	fmt.Printf("Pt1 Answer: %v\n", output)
}

func solveP2(in []int64) {
	robot := NewRobot(in)
	robot.Surface[robot.Pos] = WHITE

	robot.Start()
	for !robot.Done {
		robot.Paint()
	}
	fmt.Printf("Pt2 Answer:\n")
	robot.Surface.Draw()
}

func readInput(in io.Reader) (data []string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		data = append(data, s.Text())
	}
	data = strings.Split(data[0], ",")
	return
}

func main() {
	in := readInput(os.Stdin)
	c := Computer{}
	program := c.Compile(in)

	in1 := make([]int64, len(in))
	in2 := make([]int64, len(in))
	copy(in1, program)
	copy(in2, program)

	solveP1(in1)
	solveP2(in2)
}
