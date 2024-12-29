package day16

import (
	"fmt"
)

var dirs = map[string][]int{
	"W": {-1, 0},
	"N": {0, -1},
	"E": {1, 0},
	"S": {0, 1},
}

type pos struct {
	x int
	y int
}

type state struct {
	p    pos
	dir  string
	cost int
}

type set = map[pos]pos

func PartOne(rmap [][]string) int {
	d := "E"
	s, e, wls := startEndWalls(rmap)
	return moves(s, e, d, wls)
}

func moves(s, e pos, curDir string, wls set) int {
	return -1
}

func startEndWalls(rmap [][]string) (pos, pos, set) {
	var s pos
	var e pos
	wls := make(set)
	for y, yv := range rmap {
		for x, xv := range yv {
			if xv == "S" {
				s = pos{x: x, y: y}
			} else if xv == "E" {
				e = pos{x: x, y: y}
			} else if xv == "#" {
				wls[pos{x: x, y: y}] = pos{x: x, y: y}
			}
		}
	}
	return s, e, wls
}

func step(d string, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	return x, y
}

func printMapVisted(s, e pos, wls, vis set, h, w int) {
	for y := range h {
		for x := range w {
			if s.x == x && s.y == y {
				fmt.Print("S")
			} else if e.x == x && e.y == y {
				fmt.Print("E")
			} else if _, wok := wls[pos{x: x, y: y}]; wok {
				fmt.Print("#")
			} else if _, vok := vis[pos{x: x, y: y}]; vok {
				fmt.Print("-")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func printMap(s, e pos, wls set, h, w int) {
	for y := range h {
		for x := range w {
			if s.x == x && s.y == y {
				fmt.Print("S")
			} else if e.x == x && e.y == y {
				fmt.Print("E")
			} else if _, wok := wls[pos{x: x, y: y}]; wok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
