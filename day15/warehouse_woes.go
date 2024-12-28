package day15

import "fmt"

var dirs = map[string][]int{
	"<": {-1, 0},
	"^": {0, -1},
	">": {1, 0},
	"v": {0, 1},
}

type pos struct {
	x int
	y int
}

type set = map[pos]pos

func PartOne(wmap [][]string, dirStr string) int {
	bot, bxs, wls := botBoxesWalls(wmap)
	moves(&bot, bxs, wls, dirStr)
	return fishGPS(bxs)
}

func fishGPS(bxs set) int {
	total := 0
	for _, b := range bxs {
		total += (100 * b.y) + b.x
	}
	return total
}

func moves(bot *pos, bxs, wls set, dirStr string) {
	for _, d := range dirStr {
		var np pos
		np.x, np.y = step(string(d), bot.x, bot.y)
		_, bok := bxs[np]
		_, wok := wls[np]
		if !bok && !wok {
			*bot = np
		} else if bok {
			var nbx pos
			nbx.x, nbx.y = step(string(d), np.x, np.y)
			_, nbok := bxs[nbx]
			for nbok {
				nbx.x, nbx.y = step(string(d), nbx.x, nbx.y)
				_, nbok = bxs[nbx]
			}
			_, cwl := wls[nbx]
			if !cwl {
				bxs[nbx] = nbx
				nx, ny := step(string(d), bot.x, bot.y)
				*bot = pos{x: nx, y: ny}
				delete(bxs, *bot)
			}
		}
	}
}

func botBoxesWalls(wmap [][]string) (pos, set, set) {
	var bot pos
	bxs := make(set)
	wls := make(set)
	for y, yv := range wmap {
		for x, xv := range yv {
			if xv == "@" {
				bot = pos{x: x, y: y}
			} else if xv == "O" {
				bxs[pos{x: x, y: y}] = pos{x: x, y: y}
			} else if xv == "#" {
				wls[pos{x: x, y: y}] = pos{x: x, y: y}
			} else {
				continue
			}
		}
	}
	return bot, bxs, wls
}

func printMap(h, w int, bot pos, bxs, wls set) {
	for y := range h {
		for x := range w {
			if bot.x == x && bot.y == y {
				fmt.Print("@")
			} else if _, bok := bxs[pos{x: x, y: y}]; bok {
				fmt.Print("O")
			} else if _, wok := wls[pos{x: x, y: y}]; wok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func step(d string, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	return x, y
}
