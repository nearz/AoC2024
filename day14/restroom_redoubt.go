package day14

import (
	"fmt"
)

type pos struct {
	x int
	y int
}

type bot struct {
	p pos
	v pos
}

type grid struct {
	h int
	w int
}

const secs int = 100

var (
	finalGrid = grid{h: 103, w: 101}
	testGrid  = grid{h: 7, w: 11}
)

func PartOne(bots []bot, g grid) int {
	for i := range bots {
		moveBotSeconds(&bots[i], secs, g)
	}
	return quadCount(bots, g)
}

func quadCount(bots []bot, g grid) int {
	ym := g.h / 2
	xm := g.w / 2
	quads := map[string]int{"q1": 0, "q2": 0, "q3": 0, "q4": 0}
	for _, b := range bots {
		x, y := b.p.x, b.p.y
		if x < xm && y < ym {
			quads["q2"]++
		} else if x > xm && y < ym {
			quads["q1"]++
		} else if x < xm && y > ym {
			quads["q3"]++
		} else if x > xm && y > ym {
			quads["q4"]++
		}
	}
	qt := 1
	for _, v := range quads {
		if v != 0 {
			qt *= v
		}
	}
	return qt
}

func moveBotSeconds(b *bot, secs int, g grid) {
	dx := (b.p.x + (b.v.x * secs) + (g.w * secs)) % g.w
	dy := (b.p.y + (b.v.y * secs) + (g.h * secs)) % g.h
	b.p = pos{x: dx, y: dy}
}

func moveBot(b *bot, g grid) {
	dx := (b.p.x + b.v.x + g.w) % g.w
	dy := (b.p.y + b.v.y + g.h) % g.h
	b.p = pos{x: dx, y: dy}
}

func printBots(bots []bot, g grid) {
	bmap := make(map[pos]int)
	ym := g.h / 2
	xm := g.w / 2
	for _, b := range bots {
		bmap[b.p] += 1
	}
	for y := range g.h {
		for x := range g.w {
			if _, ok := bmap[pos{x: x, y: y}]; ok {
				fmt.Print(bmap[pos{x: x, y: y}])
			} else if x == xm {
				fmt.Print("|")
			} else if y == ym {
				fmt.Print("-")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
