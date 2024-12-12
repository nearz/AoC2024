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

const obs = "#"

type point struct {
	x int
	y int
}

type state struct {
	x   int
	y   int
	dir string
}

func PartOne(gmap [][]string) int {
	x, y, curDir := initGuardPos(gmap)
	if x == -1 {
		return -1
	}
	vis := walk(x, y, gmap, curDir)
	return len(vis)
}

func PartTwo(gmap [][]string) int {
	x, y, curDir := initGuardPos(gmap)
	if x == -1 {
		return -1
	}
	loopCount := 0
	vis := walk(x, y, gmap, curDir)
	for v := range vis {
		if x == v.x && y == v.y {
			continue
		}
		if isLoop(x, y, gmap, curDir, v) {
			loopCount++
		}
	}
	return loopCount
}

func isLoop(x, y int, gmap [][]string, curDir string, loopMark point) bool {
	loopState := make(map[state]struct{})
	inBounds := checkBounds(x, y, gmap)
	loopState[state{x: x, y: y, dir: curDir}] = struct{}{}

	for inBounds {
		cx, cy := step(curDir, x, y)
		inBounds = checkBounds(cx, cy, gmap)
		if inBounds {
			if gmap[cy][cx] == obs || (cy == loopMark.y && cx == loopMark.x) {
				curDir = rotate(curDir)
			} else {
				x = cx
				y = cy
			}
			newState := state{x: x, y: y, dir: curDir}
			if _, ex := loopState[newState]; ex {
				return true
			}
			loopState[newState] = struct{}{}
		}

	}
	return false
}

func walk(x, y int, gmap [][]string, curDir string) map[point]struct{} {
	vis := make(map[point]struct{})
	inBounds := checkBounds(x, y, gmap)
	for inBounds {
		cx, cy := step(curDir, x, y)
		inBounds = checkBounds(cx, cy, gmap)
		if inBounds {
			if gmap[cy][cx] != obs {
				x = cx
				y = cy
			} else {
				curDir = rotate(curDir)
			}
			vis[point{x: x, y: y}] = struct{}{}
		}
	}
	return vis
}

func initGuardPos(gmap [][]string) (int, int, string) {
	for y, row := range gmap {
		for x, col := range row {
			for k := range guardDirs {
				if k == col {
					curDir := guardDirs[k]
					return x, y, curDir
				}
			}
		}
	}
	return -1, -1, ""
}

func rotate(curDir string) string {
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
	return curDir
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
