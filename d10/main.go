package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

type Asteroid [2]int
type AsteroidField map[Asteroid]struct{}
type LOSMap map[float64][]Asteroid

func (a Asteroid) AngleTo(dest Asteroid) float64 {
	x := float64(dest[0] - a[0])
	y := float64(dest[1] - a[1])
	return math.Atan2(x, y)
}

func (a Asteroid) DistanceTo(dest Asteroid) float64 {
	x := float64(dest[0] - a[0])
	y := float64(dest[1] - a[1])
	return math.Sqrt(x*x + y*y)
}

func (a Asteroid) FindNearest(list []Asteroid) (closest Asteroid) {
	for i, asteroid := range list {
		if i == 0 || a.DistanceTo(asteroid) < a.DistanceTo(closest) {
			closest = asteroid
		}
	}
	return
}

// Count the asteroids in the line of sight
func (a Asteroid) CountLOS(af AsteroidField) int {
	return len(a.LOSMap(af))
}

func (a Asteroid) LOSMap(af AsteroidField) LOSMap {
	losMap := make(LOSMap)
	for dest, _ := range af {
		if a != dest {
			angle := a.AngleTo(dest)
			losMap[angle] = append(losMap[angle], dest)
		}
	}
	return losMap
}

func NewAsteroidField(rows []string) AsteroidField {
	af := AsteroidField{}
	for y, row := range rows {
		for x, chr := range row {
			if chr == '#' {
				af[Asteroid{x, y}] = struct{}{}
			}
		}
	}
	return af
}

func (af AsteroidField) Destroy(a Asteroid) {
	delete(af, a)
}

func (losm LOSMap) Angles() []float64 {
	angles := make([]float64, len(losm))
	i := 0
	for angle := range losm {
		angles[i] = angle
		i++
	}
	return angles
}

func bestStation(af AsteroidField) (Asteroid, int) {
	var best Asteroid
	bestLOS := 0
	for src, _ := range af {
		LOSCount := src.CountLOS(af)
		if LOSCount > bestLOS {
			bestLOS = LOSCount
			best = src
		}
	}
	return best, bestLOS
}

func pewPew(af AsteroidField, best Asteroid, nth int) Asteroid {
	destroyed := 0
	losMap := best.LOSMap(af)
	angles := losMap.Angles()
	sort.Slice(angles, func(i, j int) bool { return angles[i] > angles[j] }) //reverse
	for {
		losMap := best.LOSMap(af)
		for _, angle := range angles {
			line := losMap[angle]
			if len(line) == 0 {
				continue
			}
			nearest := best.FindNearest(line)
			af.Destroy(nearest)
			destroyed++
			if destroyed == nth {
				return nearest
			}
		}
	}
}

func solveP1(af AsteroidField) {
	best, bestLOS := bestStation(af)
	fmt.Printf("Pt1 Answer: best:%v count:%d\n", best, bestLOS)
}

func solveP2(af AsteroidField) {
	station, _ := bestStation(af)
	nth := pewPew(af, station, 200)
	fmt.Printf("Pt2 Answer: %v\n", nth[0]*100+nth[1])
}

func readInput(in io.Reader) (data []string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		data = append(data, s.Text())
	}
	return
}

func main() {
	in := readInput(os.Stdin)
	af := NewAsteroidField(in)
	solveP1(af)
	solveP2(af)
}
