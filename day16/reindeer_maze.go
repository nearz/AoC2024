package day16

import (
	"container/heap"
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
	ss := state{p: s, dir: curDir, cost: 0}
	vis := make(map[string]struct{})
	q := StateHeap{ss}
	for len(q) > 0 {
		cur := heap.Pop(&q).(state)
		key := fmt.Sprintf("%d,%d,%s", cur.p.x, cur.p.y, cur.dir)
		if _, vok := vis[key]; vok {
			continue
		}
		vis[key] = struct{}{}
		if cur.p == e {
			return cur.cost
		}
		for d := range dirs {
			x, y := step(d, cur.p.x, cur.p.y)
			_, wok := wls[pos{x: x, y: y}]
			if !wok {
				nc := cur.cost + 1
				if d != cur.dir {
					nc += 1000
				}
				ns := state{p: pos{x: x, y: y}, dir: d, cost: nc}
				heap.Push(&q, ns)
			}
		}

	}
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
