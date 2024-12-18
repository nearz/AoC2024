package day10

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

func PartOne(tmap [][]int) int {
	ths := trailheads(tmap)
	scores := make(map[pos]int)
	for _, t := range ths {
		vis := make(map[pos]struct{})
		v := tmap[t.y][t.x]
		scores[t] = walkTrail(v, t, vis, tmap)
	}
	total := 0
	for s := range scores {
		total += scores[s]
	}
	return total
}

func walkTrail(n int, p pos, vis map[pos]struct{}, tmap [][]int) int {
	count := 0
	if n == 9 {
		if _, ex := vis[p]; ex {
			return 0
		}
		vis[p] = struct{}{}
		return 1
	}
	for k := range dirs {
		v := peek(k, p.x, p.y, tmap)
		if v-tmap[p.y][p.x] == 1 {
			var nextPos pos
			nextPos.x, nextPos.y = step(k, p.x, p.y)
			count += walkTrail(v, nextPos, vis, tmap)

		}
	}
	return count
}

func trailheads(tmap [][]int) []pos {
	ths := []pos{}
	for y, yv := range tmap {
		for x, xv := range yv {
			if xv == 0 {
				ths = append(ths, pos{x: x, y: y})
			}
		}
	}
	return ths
}

func peek(d string, x, y int, tmap [][]int) int {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	if !checkBounds(x, y, tmap) {
		return -1
	}
	return tmap[y][x]
}

func step(d string, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	return x, y
}

func checkBounds(x, y int, m [][]int) bool {
	ux, uy := len(m[0]), len(m)
	b := (-1 < x && x < ux) && (-1 < y && y < uy)
	return b
}
