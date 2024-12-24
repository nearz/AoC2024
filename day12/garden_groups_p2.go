package day12

type edge struct {
	p  pos
	nb pos
}

func PartTwo(garden [][]string) int {
	avl := available(garden)
	total := 0
	for a := range avl {
		region := make(map[pos]struct{})
		rgnVal := garden[a.y][a.x]
		getRegion(a, rgnVal, avl, region, garden)
		// fmt.Printf("L: %s | Area: %d", rgnVal, len(region))
		// fmt.Printf(" Sides: %d\n", sides(region))
		total += len(region) * sides(region)

	}
	return total
}

func sides(region map[pos]struct{}) int {
	sides := 0
	edges := make(map[edge]struct{})
	for p := range region {
		for d := range dirs {
			x, y := step(d, p.x, p.y)
			var np pos
			np.x, np.y = x, y
			if _, ok := region[np]; !ok {
				e := edge{p: p, nb: np}
				edges[e] = struct{}{}
			}

		}
	}
	for pe := range edges {
		if pe.p.x == pe.nb.x {
			npp := stepLeft(pe.p)
			nnb := stepLeft(pe.nb)
			ne := edge{p: npp, nb: nnb}
			if _, ok := edges[ne]; !ok {
				sides += 1
			}
		} else {
			npp := stepDown(pe.p)
			nnb := stepDown(pe.nb)
			ne := edge{p: npp, nb: nnb}
			if _, ok := edges[ne]; !ok {
				sides += 1
			}
		}
	}
	return sides
}

func stepDown(p pos) pos {
	x, y := step("d", p.x, p.y)
	var np pos
	np.x, np.y = x, y
	return np
}

func stepLeft(p pos) pos {
	x, y := step("l", p.x, p.y)
	var np pos
	np.x, np.y = x, y
	return np
}
