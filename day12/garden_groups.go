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

func PartOne(garden [][]string) int {
	avl := available(garden)
	total := 0
	for a := range avl {
		region := make(map[pos]struct{})
		rgnVal := garden[a.y][a.x]
		getRegion(a, rgnVal, avl, region, garden)
		perim := 0
		for r := range region {
			perim += perimeter(r, garden)
		}
		total += len(region) * perim
	}
	return total
}

func getRegion(p pos, rgnVal string, avl map[pos]struct{}, region map[pos]struct{}, garden [][]string) {
	delete(avl, p)
	region[p] = struct{}{}
	for d := range dirs {
		x, y := step(d, p.x, p.y)
		if !checkBounds(x, y, garden) {
			continue
		}
		var np pos
		np.x, np.y = x, y
		nv := garden[np.y][np.x]
		_, ok := avl[np]
		if nv == rgnVal && ok {
			getRegion(np, rgnVal, avl, region, garden)
		}
	}
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
