package day12

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

type plot struct {
	perm int
	area int
}

func PartOne(garden [][]string) int {
	avl := available(garden)
	res := make(map[pos]plot)
	for a := range avl {
		area, pmtr := walkPlot(0, a, avl, garden)
		plt := plot{perm: pmtr, area: area}
		res[a] = plt
	}
	total := 0
	for r := range res {
		plotTotal := res[r].area * res[r].perm
		total += plotTotal
	}
	return total
}

func walkPlot(t int, p pos, avl map[pos]struct{}, garden [][]string) (int, int) {
	delete(avl, p)
	area := 0
	pmtr := 0
	if t == 0 {
		area++
		pmtr += perimeter(p, garden)
	}
	for d := range dirs {
		v := peek(d, p.x, p.y, garden)
		var np pos
		np.x, np.y = step(d, p.x, p.y)
		_, ok := avl[np]
		if v == garden[p.y][p.x] && ok {
			area++
			pmtr += perimeter(np, garden)
			a, p := walkPlot(t+1, np, avl, garden)
			area += a
			pmtr += p
		}
	}
	return area, pmtr
}

func perimeter(p pos, garden [][]string) int {
	pmtr := 0
	for d := range dirs {
		v := peek(d, p.x, p.y, garden)
		if v != garden[p.y][p.x] {
			pmtr++
		}
	}
	return pmtr
}

func available(garden [][]string) map[pos]struct{} {
	avl := make(map[pos]struct{})
	for i, ig := range garden {
		for j := range ig {
			p := pos{x: j, y: i}
			avl[p] = struct{}{}
		}
	}
	return avl
}

func peek(d string, x, y int, tmap [][]string) string {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	if !checkBounds(x, y, tmap) {
		return "-1"
	}
	return tmap[y][x]
}

func step(d string, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	return x, y
}

func checkBounds(x, y int, m [][]string) bool {
	ux, uy := len(m[0]), len(m)
	b := (-1 < x && x < ux) && (-1 < y && y < uy)
	return b
}
