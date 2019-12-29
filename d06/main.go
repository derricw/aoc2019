package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Object struct {
	Name       string
	Parent     *Object
	Satellites []*Object
}

func (o *Object) CountOrbits() int {
	if o.Parent == nil {
		return 0
	} else {
		return 1 + o.Parent.CountOrbits()
	}
}

func NewObject(name string) *Object {
	return &Object{
		Name:       name,
		Parent:     nil,
		Satellites: make([]*Object, 0),
	}
}

type OrbitMap struct {
	Objects map[string]*Object
	COM     *Object
}

func (om *OrbitMap) AddOrbit(orbit string) {
	orb := strings.Split(orbit, ")")
	parent, sat := orb[0], orb[1]
	var pobj, sobj *Object
	if obj, ok := om.Objects[parent]; ok {
		pobj = obj
	} else {
		pobj = NewObject(parent)
	}
	if obj, ok := om.Objects[sat]; ok {
		sobj = obj
	} else {
		sobj = NewObject(sat)
	}
	pobj.Satellites = append(pobj.Satellites, sobj)
	sobj.Parent = pobj
	om.Objects[parent] = pobj
	om.Objects[sat] = sobj
	if parent == "COM" {
		om.COM = pobj
	}
}

func (om *OrbitMap) TotalOrbits() int {
	total := 0
	for _, obj := range om.Objects {
		total += obj.CountOrbits()
	}
	return total
}

func (om *OrbitMap) CalculateDistance(src, dest string) int {
	if src == dest {
		return 0
	}
	type Visit struct {
		Name     string
		Distance int
	}
	toVisit := make([]*Visit, 0)
	visited := make(map[string]struct{})
	sobj := om.Objects[src]
	toVisit = append(toVisit, &Visit{Name: sobj.Parent.Name, Distance: 1})
	for _, obj := range sobj.Satellites {
		toVisit = append(toVisit, &Visit{Name: obj.Name, Distance: 1})
	}

	var current *Visit
	for {
		if len(toVisit) == 0 {
			panic("failed to find route")
		}
		current, toVisit = toVisit[0], toVisit[1:] // frontpop!
		if _, ok := visited[current.Name]; ok {
			continue
		}
		if current.Name == dest {
			return current.Distance
		}
		currObj := om.Objects[current.Name]
		if currObj.Parent != nil {
			toVisit = append(toVisit, &Visit{
				Name:     currObj.Parent.Name,
				Distance: current.Distance + 1,
			})
		}
		for _, obj := range currObj.Satellites {
			toVisit = append(toVisit, &Visit{
				Name:     obj.Name,
				Distance: current.Distance + 1,
			})
		}
		visited[current.Name] = struct{}{}
	}
}

func NewOrbitMap(orbitList []string) *OrbitMap {
	objMap := make(map[string]*Object)
	orbMap := &OrbitMap{Objects: objMap}
	for _, orbit := range orbitList {
		orbMap.AddOrbit(orbit)
	}
	return orbMap
}

func solveP1(om *OrbitMap) {
	orbits := om.TotalOrbits()
	fmt.Printf("Pt1 Answer: %d\n", orbits)
}

func solveP2(om *OrbitMap) {
	youParent := om.Objects["YOU"].Parent.Name
	sanParent := om.Objects["SAN"].Parent.Name
	distance := om.CalculateDistance(youParent, sanParent)
	fmt.Printf("Pt2 Answer: %d\n", distance)
}

func readInput(in io.Reader) (data []string) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		data = append(data, s.Text())
	}
	//data = strings.Split(data[0], ",")
	return
}

func main() {
	in := readInput(os.Stdin)
	om := NewOrbitMap(in)
	solveP1(om)
	solveP2(om)
}
