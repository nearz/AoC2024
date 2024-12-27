package day14

import "fmt"

var dirs = map[string][]int{
	"l": {-1, 0},
	"u": {0, -1},
	"r": {1, 0},
	"d": {0, 1},
}

// find the largest cluster of neigboring bots to find the christmas tree
func PartTwo(bots []bot, g grid) {
	bigCluster := 0.0
	var start pos
	var secs int
	tree := make(map[pos]struct{})
	for s := range 10000 {
		bset := botPosSet(bots)
		bl := len(bset)
		for b := range bset {
			vis := cluster(b, bset)
			vratio := float64(len(vis)) / float64(bl)
			if vratio > bigCluster {
				bigCluster = vratio
				start = b
				secs = s
				tree = vis
			}
			for v := range vis {
				delete(bset, v)
			}
		}
		for i := range bots {
			moveBot(&bots[i], g)
		}
	}
	fmt.Printf("start: %v, bigc: %v, secs: %d, tree: %v\n", start, bigCluster, secs, tree)
}

func cluster(p pos, bset map[pos]struct{}) map[pos]struct{} {
	vis := make(map[pos]struct{})
	q := []pos{p}
	vis[p] = struct{}{}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for d := range dirs {
			var np pos
			np.x, np.y = step(d, cur.x, cur.y)
			_, bok := bset[np]
			_, vok := vis[np]
			if bok && !vok {
				vis[np] = struct{}{}
				q = append(q, np)
			}
		}
	}
	return vis
}

func botPosSet(bots []bot) map[pos]struct{} {
	bset := make(map[pos]struct{})
	for _, b := range bots {
		bset[b.p] = struct{}{}
	}
	return bset
}

func step(d string, x, y int) (int, int) {
	dx, dy := dirs[d][0], dirs[d][1]
	x = x + dx
	y = y + dy
	return x, y
}
