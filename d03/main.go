package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vec2D [2]int

var Origin Vec2D = Vec2D{0, 0}

// Directions
var D map[string]Vec2D = map[string]Vec2D{
	"U": Vec2D{0, 1},
	"D": Vec2D{0, -1},
	"L": Vec2D{-1, 0},
	"R": Vec2D{1, 0},
}

// Dot product
func (v Vec2D) Dot(u Vec2D) Vec2D {
	return Vec2D{v[0] * u[0], v[1] * u[1]}
}

// Translate
func (v Vec2D) Trans(u Vec2D) Vec2D {
	return Vec2D{v[0] + u[0], v[1] + u[1]}
}

func NewVec2DFromString(vecString string) Vec2D {
	magStr := strings.Split(vecString, ",")
	magx, _ := strconv.Atoi(magStr[0])
	magy, _ := strconv.Atoi(magStr[1])
	return Vec2D{magx, magy}
}

// so cool how math only works with floats
func ManDist(p0, p1 Vec2D) int {
	return int(math.Abs(float64(p1[0]-p0[0])) +
		math.Abs(float64(p1[1]-p0[1])))
}

// Lets save the points of our wire as a map for fast lookup
type Wire map[string]struct{}

var Used struct{}

func NewWireFromWirePath(wp WirePath) Wire {
	w := Wire{
		"0,0": Used,
	}
	for _, pos := range wp {
		w[fmt.Sprintf("%d,%d", pos[0], pos[1])] = Used
	}
	return w
}

func NewWireFromPath(path []string) Wire {
	wp := NewWirePathFromPath(path)
	return NewWireFromWirePath(wp)
}

func FindIntersections(w0, w1 Wire) (cross []Vec2D) {
	for p, _ := range w0 {
		if p != "0,0" {
			_, found := w1[p]
			if found {
				cross = append(cross, NewVec2DFromString(p))
			}
		}
	}
	return cross
}

func DistanceToClosest(crosses []Vec2D) (minDist int) {
	for i := 0; i < len(crosses); i++ {
		manDist := ManDist(Origin, crosses[i])
		if i == 0 || manDist < minDist {
			minDist = manDist
		}
	}
	return
}

type WirePath []Vec2D

func NewWirePathFromPath(path []string) (wp WirePath) {
	pos := Origin
	for _, p := range path {
		dir := string(p[0])
		mag, _ := strconv.Atoi(p[1:len(p)])
		for i := 0; i < mag; i++ {
			pos = pos.Trans(D[dir])
			wp = append(wp, pos)
		}
	}
	return wp
}

func FindShortestPath(w0, w1 WirePath, crosses []Vec2D) int {
	var minDist int
	for i, cross := range crosses {
		var dist0 int = 0
		var dist1 int = 0
		for _, p := range w0 {
			dist0++
			if p[0] != cross[0] || p[1] != cross[1] {
			} else {
				break
			}
		}
		for _, p := range w1 {
			dist1++
			if p[0] != cross[0] || p[1] != cross[1] {
			} else {
				break
			}
		}
		total := dist0 + dist1
		if i == 0 || total < minDist {
			minDist = total
		}
	}
	return minDist
}

func solve_p1(in [][]string) {
	w0 := NewWireFromPath(in[0])
	w1 := NewWireFromPath(in[1])
	crosses := FindIntersections(w0, w1)
	closest := DistanceToClosest(crosses)
	fmt.Printf("Pt1 Answer: %d\n", closest)
}

func solve_p2(in [][]string) {
	wp0 := NewWirePathFromPath(in[0])
	wp1 := NewWirePathFromPath(in[1])
	w0 := NewWireFromWirePath(wp0)
	w1 := NewWireFromWirePath(wp1)

	crosses := FindIntersections(w0, w1)
	shortestPath := FindShortestPath(wp0, wp1, crosses)

	fmt.Printf("Pt2 Answer: %d\n", shortestPath)
}

func readInput(in io.Reader) (paths [][]string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		paths = append(paths, strings.Split(s.Text(), ","))
	}
	return
}

func main() {
	in := readInput(os.Stdin)

	in1 := make([][]string, len(in))
	in2 := make([][]string, len(in))
	copy(in1, in)
	copy(in2, in)

	solve_p1(in1)
	solve_p2(in2)
}
