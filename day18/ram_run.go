package day18

import (
	"fmt"
)

var dirs = map[string][]int{
	"l": {-1, 0},
	"u": {0, -1},
	"r": {1, 0},
	"d": {0, 1},
}

type pos struct {
	x int
	y int
}

type grid struct {
	h int
	w int
	s pos
	e pos
}

type nodeState struct {
	p     pos
	steps int
}

type posSet map[pos]struct{}

var (
	testGrid  = grid{h: 7, w: 7, s: pos{x: 0, y: 0}, e: pos{x: 6, y: 6}}
	finalGrid = grid{h: 71, w: 71, s: pos{x: 0, y: 0}, e: pos{x: 70, y: 70}}
)

const (
	tbLen = 12
	bLen  = 1024
)

func PartTwo(pslc []pos, g grid) pos {
	var blockPos pos
	l := len(pslc)
	for i := 1; i < l; i++ {
		byts := convToPset(pslc, i)
		ns := move(byts, g)
		if ns.steps == 0 {
			blockPos = pslc[i-1]
			break
		}
	}
	return blockPos
}

func PartOne(pslc []pos, g grid, byteLen int) int {
	byts := convToPset(pslc, byteLen)
	return move(byts, g).steps
}

func move(byts posSet, g grid) nodeState {
	var fn nodeState
	vis := make(posSet)
	s := nodeState{p: g.s, steps: 0}
	q := []nodeState{s}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if _, vok := vis[cur.p]; vok {
			continue
		}
		vis[cur.p] = struct{}{}
		if cur.p == g.e {
			fn = cur
			break
		}
		for d := range dirs {
			x, y := step(d, cur.p.x, cur.p.y)
			inb := inBounds(x, y, g.h, g.w)
			np := pos{x: x, y: y}
			_, bok := byts[np]
			if !bok && inb {
				var ns nodeState
				ns.p = np
				ns.steps = cur.steps + 1
				q = append(q, ns)
			}
		}

	}
	return fn
}

func step(d string, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	return x, y
}

func inBounds(x, y, h, w int) bool {
	b := (-1 < x && x < w) && (-1 < y && y < h)
	return b
}

func printGrid(byts, vis posSet, g grid) {
	for y := range g.h {
		for x := range g.w {
			p := pos{x: x, y: y}
			if _, ok := byts[p]; ok {
				fmt.Print("#")
			} else if _, vok := vis[p]; vok {
				fmt.Print("O")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func convToPset(pslc []pos, n int) posSet {
	pset := make(posSet)
	for i := range n {
		pset[pslc[i]] = struct{}{}
	}
	return pset
}
