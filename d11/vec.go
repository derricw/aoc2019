package main

import (
	"strconv"
	"strings"
)

type Vec2D [2]int

var Origin Vec2D = Vec2D{0, 0}

var D map[string]Vec2D = map[string]Vec2D{
	"U": Vec2D{0, 1},
	"D": Vec2D{0, -1},
	"L": Vec2D{-1, 0},
	"R": Vec2D{1, 0},
}

func (v Vec2D) Trans(u Vec2D) Vec2D {
	return Vec2D{v[0] + u[0], v[1] + u[1]}
}

func NewVec2DFromString(vecString string) Vec2D {
	magStr := strings.Split(vecString, ",")
	magx, _ := strconv.Atoi(magStr[0])
	magy, _ := strconv.Atoi(magStr[1])
	return Vec2D{magx, magy}
}
