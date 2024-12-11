package day6

var dirs = map[string][]int{
	"l": {-1, 0},
	"u": {0, -1},
	"r": {1, 0},
	"d": {0, 1},
}

var guardDirs = map[string]string{
	"^": "u",
	">": "r",
	"v": "d",
	"<": "l",
}

var curDir = "u"

const (
	obs  = "#"
	mark = "X"
)

func PartOne(gmap [][]string) int {
	for y, row := range gmap {
		for x, col := range row {
			for k := range guardDirs {
				if k == col {
					curDir = guardDirs[k]
					walk(x, y, gmap)
					break
				}
			}
		}
	}
	return countPos(gmap, mark)
}

func PartTwo(gmap [][]string) int {
	return -1
}

func walk(x, y int, gmap [][]string) {
	inBounds := checkBounds(x, y, gmap)
	for inBounds {
		cx, cy := step(curDir, x, y)
		inBounds = checkBounds(cx, cy, gmap)
		if inBounds {
			if gmap[cy][cx] != obs {
				gmap[cy][cx] = mark
				x = cx
				y = cy
			} else {
				rotate()
			}
		}
	}
}

func countPos(gmap [][]string, item string) int {
	count := 0
	for _, row := range gmap {
		for _, col := range row {
			if col == item {
				count++
			}
		}
	}
	return count
}

func rotate() {
	switch curDir {
	case "l":
		curDir = "u"
	case "u":
		curDir = "r"
	case "r":
		curDir = "d"
	case "d":
		curDir = "l"
	}
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
