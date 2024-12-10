package day4

var dirs = map[string][]int{
	"l":  {-1, 0},
	"ul": {-1, -1},
	"u":  {0, -1},
	"ur": {1, -1},
	"r":  {1, 0},
	"dr": {1, 1},
	"d":  {0, 1},
	"dl": {-1, 1},
}

const find = "XMAS"

func PartOne(cs [][]string) int {
	finalCount := 0
	for i := range cs {
		for j := range len(cs[i]) {
			checkDirs := cs[i][j] == string(find[0])
			if checkDirs {
				finalCount += allDirs(j, i, cs)
			}
		}
	}
	return finalCount
}

func PartTwo(cs [][]string) int {
	finalCount := 0
	for i := range cs {
		for j := range len(cs[i]) {
			checkXmas := cs[i][j] == "A"
			if checkXmas {
				check := isXmas(j, i, cs)
				if check {
					finalCount++
				}
			}
		}
	}
	return finalCount
}

func allDirs(x, y int, cs [][]string) int {
	lf := len(find)
	foundCount := 0
	for k := range dirs {
		found := true
		for i := 1; i < lf; i++ {
			char := string(find[i])
			cx, cy := move(k, i, x, y)
			inBounds := checkBounds(cx, cy, cs)
			if !inBounds {
				found = false
				break
			}
			isChar := cs[cy][cx] == char
			if !isChar {
				found = false
				break
			}
		}
		if found {
			foundCount++
		}
	}
	return foundCount
}

func isXmas(x, y int, cs [][]string) bool {
	mas := "MAS"
	rmas := "SAM"
	legOne := oneOffChar(x, y, cs, "ul") + "A" + oneOffChar(x, y, cs, "dr")
	legTwo := oneOffChar(x, y, cs, "dl") + "A" + oneOffChar(x, y, cs, "ur")
	return (legOne == mas || legOne == rmas) && (legTwo == mas || legTwo == rmas)
}

func oneOffChar(x, y int, cs [][]string, dir string) string {
	cx, cy := move(dir, 1, x, y)
	inBounds := checkBounds(cx, cy, cs)
	if !inBounds {
		return ""
	}
	return cs[cy][cx]
}

func move(d string, dis, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + (dis * dx)
	y = y + (dis * dy)
	return x, y
}

func checkBounds(x, y int, m [][]string) bool {
	ux, uy := len(m[0]), len(m)
	b := (-1 < x && x < ux) && (-1 < y && y < uy)
	return b
}
