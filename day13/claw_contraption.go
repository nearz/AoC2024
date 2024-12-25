package day13

type button struct {
	x    int
	y    int
	cost int
}

type game struct {
	a button
	b button
	x int
	y int
}

const (
	maxClicks   = 100
	prizeOffset = 10_000_000_000_000
)

func PartOne(games []game) int {
	total := 0
	for _, g := range games {
		total += mincost(g.a.x, g.a.y, g.b.x, g.b.y, g.x, g.y)
	}
	return total
}

func PartTwo(games []game) int {
	total := 0
	for _, g := range games {
		gxo := g.x + prizeOffset
		gyo := g.y + prizeOffset
		total += mincost(g.a.x, g.a.y, g.b.x, g.b.y, gxo, gyo)
	}
	return total
}

func mincost(ax, ay, bx, by, gx, gy int) int {
	a, ar := divmod(by*gx-bx*gy, by*ax-bx*ay)
	b, br := divmod(gx-ax*a, bx)
	if ar != 0 || br != 0 {
		return 0
	}
	return a*3 + b
}

func divmod(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
