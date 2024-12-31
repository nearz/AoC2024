package day16

import (
	"container/heap"
	"fmt"
	"math"
)

type track struct {
	s     state
	track []state
}

func PartTwo(rmap [][]string) int {
	d := "E"
	s, e, wls := startEndWalls(rmap)
	return endToStart(s, e, d, wls)
}

func endToStart(s, e pos, curDir string, wls set) int {
	minCost := math.MaxInt
	ss := state{p: s, dir: curDir, cost: 0}
	vis := make(map[string]int)
	paths := [][]state{}
	t := track{
		s:     ss,
		track: []state{ss},
	}
	q := TrackHeap{t}
	for q.Len() > 0 {
		cur := heap.Pop(&q).(track)
		if cur.s.cost > minCost {
			break
		}
		vkey := fmt.Sprintf("%d,%d,%s", cur.s.p.x, cur.s.p.y, cur.s.dir)
		if c, vok := vis[vkey]; vok && c < cur.s.cost {
			continue
		}
		vis[vkey] = cur.s.cost
		if cur.s.p == e {
			minCost = cur.s.cost
			paths = append(paths, cur.track)
		}

		for d := range dirs {
			x, y := step(d, cur.s.p.x, cur.s.p.y)
			if _, wok := wls[pos{x: x, y: y}]; !wok {
				nc := cur.s.cost + 1
				if d != cur.s.dir {
					nc += 1000
				}
				ns := state{p: pos{x: x, y: y}, dir: d, cost: nc}
				var nt track
				nt.s = ns
				nt.track = append([]state{}, cur.track...)
				nt.track = append(nt.track, ns)
				heap.Push(&q, nt)
			}
		}
	}
	finalNodes := make(map[pos]struct{})
	for _, p := range paths {
		for _, s := range p {
			finalNodes[s.p] = struct{}{}
		}
	}
	return len(finalNodes)
}

func printMap(ns map[pos]struct{}, wls set, h, w int) {
	for y := range h {
		for x := range w {
			p := pos{x: x, y: y}
			if _, nok := ns[p]; nok {
				fmt.Print("O")
			} else if _, wok := wls[p]; wok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
