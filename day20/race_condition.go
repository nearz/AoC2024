package day20

import (
	"strconv"
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

type posSet map[pos]struct{}

func PartOne(rcmap [][]string) int {
	s, e, wls, pth := startEndWallsPath(rcmap)
	cwls := candidateWalls(s, e, wls, pth)
	plst := pathList(s, e, pth)
	timeMap := make(map[int]int)
	for c := range cwls {
		adjPair := adjacentPair(c, wls)
		one := -1
		two := -1
		for i, p := range plst {
			if _, aok := adjPair[p]; aok && one == -1 {
				one = i
			} else if _, aok := adjPair[p]; aok && two == -1 {
				two = i
			}
		}
		save := two - one - 2
		if _, tok := timeMap[save]; tok {
			timeMap[save] += 1
		} else {
			timeMap[save] = 1
		}
	}
	total := 0
	for k, v := range timeMap {
		if k >= 100 {
			total += v
		}
	}
	return total
}

func adjacentPair(p pos, wls posSet) posSet {
	adjSet := make(posSet)
	for d := range dirs {
		np := stepPos(d, p)
		if _, wok := wls[np]; !wok {
			adjSet[np] = struct{}{}
		}
	}
	if len(adjSet) > 2 {
		rmmap := make(map[string][]pos)
		rAdjSet := make(posSet)
		for a := range adjSet {
			x := strconv.Itoa(a.x) + "x"
			y := strconv.Itoa(a.y) + "y"
			rmmap[x] = append(rmmap[x], a)
			rmmap[y] = append(rmmap[y], a)
		}
		for _, v := range rmmap {
			if len(v) > 1 {
				for _, k := range v {
					rAdjSet[k] = struct{}{}
				}
			}
		}
		adjSet = rAdjSet
	}
	return adjSet
}

func pathList(s, e pos, pth posSet) []pos {
	type state struct {
		p     pos
		steps []pos
	}
	vis := make(posSet)
	vis[s] = struct{}{}
	var ss state
	ss.p = s
	ss.steps = append(ss.steps, s)
	q := []state{ss}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.p == e {
			return cur.steps
		}
		for d := range dirs {
			np := stepPos(d, cur.p)
			_, pok := pth[np]
			_, vok := vis[np]
			if pok && !vok {
				vis[np] = struct{}{}
				var ns state
				ns.p = np
				ns.steps = append(ns.steps, cur.steps...)
				ns.steps = append(ns.steps, np)
				q = append(q, ns)
			}
		}
	}
	return []pos{}
}

func candidateWalls(s, e pos, wls, pth posSet) posSet {
	cwls := make(posSet)
	for w := range wls {
		_, up := pth[stepPos("u", w)]
		sup := stepPos("u", w) == s
		eup := stepPos("u", w) == e
		_, dn := pth[stepPos("d", w)]
		sdn := stepPos("d", w) == s
		edn := stepPos("d", w) == e
		_, rt := pth[stepPos("r", w)]
		srt := stepPos("r", w) == s
		ert := stepPos("r", w) == e
		_, lt := pth[stepPos("l", w)]
		slt := stepPos("l", w) == s
		elt := stepPos("l", w) == e
		if (up || sup || eup) && (dn || sdn || edn) {
			cwls[w] = struct{}{}
		} else if (rt || srt || ert) && (lt || slt || elt) {
			cwls[w] = struct{}{}
		}
	}
	return cwls
}

func startEndWallsPath(rcmap [][]string) (pos, pos, posSet, posSet) {
	var s pos
	var e pos
	wls := make(posSet)
	pth := make(posSet)
	for y, yv := range rcmap {
		for x, xv := range yv {
			p := pos{x: x, y: y}
			if xv == "S" {
				s = p
				pth[p] = struct{}{}
			} else if xv == "E" {
				e = p
				pth[p] = struct{}{}
			} else if xv == "#" {
				wls[p] = struct{}{}
			} else if xv == "." {
				pth[p] = struct{}{}
			}
		}
	}
	return s, e, wls, pth
}

func stepPos(d string, p pos) pos {
	var np pos
	dx, dy := dirs[d][0], dirs[d][1]
	np.x = p.x + dx
	np.y = p.y + dy
	return np
}
