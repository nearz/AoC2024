package day8

type point struct {
	x int
	y int
}

func PartOne(amap [][]string) int {
	apos := positions(amap)
	return countAnode(apos, amap)
}

func PartTwo(amap [][]string) int {
	apos := positions(amap)
	return countHarmonics(apos, amap)
}

func countHarmonics(apos map[string][]point, amap [][]string) int {
	hset := make(map[point]struct{})
	for k := range apos {
		ps := apos[k]
		l := len(ps)
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if ps[i] != ps[j] {
					hset[point{x: ps[i].x, y: ps[i].y}] = struct{}{}
					d := dis(ps[i], ps[j])
					y := ps[i].y + d.y
					x := ps[i].x + d.x
					for checkBounds(x, y, amap) {
						hset[point{x: x, y: y}] = struct{}{}
						y = y + d.y
						x = x + d.x
					}
				}
			}
		}
	}
	return len(hset)
}

func countAnode(apos map[string][]point, amap [][]string) int {
	anset := make(map[point]struct{})
	for k := range apos {
		ps := apos[k]
		l := len(ps)
		for i := 0; i < l; i++ {
			for j := 0; j < l; j++ {
				if ps[i] != ps[j] {
					d := dis(ps[i], ps[j])
					y := ps[i].y + d.y
					x := ps[i].x + d.x
					if checkBounds(x, y, amap) {
						anset[point{x: x, y: y}] = struct{}{}
					}
				}
			}
		}
	}
	return len(anset)
}

func positions(amap [][]string) map[string][]point {
	apos := make(map[string][]point)
	for i, iv := range amap {
		for j, jv := range iv {
			if jv != "." {
				p := point{x: j, y: i}
				apos[jv] = append(apos[jv], p)
			}
		}
	}
	return apos
}

func checkBounds(x, y int, m [][]string) bool {
	ux, uy := len(m[0]), len(m)
	b := (-1 < x && x < ux) && (-1 < y && y < uy)
	return b
}

func dis(a point, b point) point {
	np := point{
		x: a.x - b.x,
		y: a.y - b.y,
	}
	return np
}

// func anodeMap(apos map[string][]point) map[point][]point {
// 	anodeMap := make(map[point][]point)
// 	for k := range apos {
// 		ps := apos[k]
// 		l := len(ps)
// 		for i := 0; i < l; i++ {
// 			for j := 0; j < l; j++ {
// 				if ps[i] != ps[j] {
// 					d := dis(ps[i], ps[j])
// 					anodeMap[ps[i]] = append(anodeMap[ps[i]], d)
//
// 				}
// 			}
// 		}
// 	}
// 	return anodeMap
// }

// func placeAnode(anpos map[point][]point, amap [][]string) int {
// 	anset := make(map[point]struct{})
// 	for k := range anpos {
// 		ps := anpos[k]
// 		for _, p := range ps {
// 			y := k.y + p.y
// 			x := k.x + p.x
// 			if checkBounds(x, y, amap) {
// 				anset[point{x: x, y: y}] = struct{}{}
// 			}
// 		}
// 	}
// 	return len(anset)
// }
