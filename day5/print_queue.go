package day5

func PartOne(rules [][]int, pqs [][]int) int {
	pm := priorityMap(rules)
	cqs, _ := sortedQueues(pm, pqs)
	total := totalOfMiddles(cqs)
	return total
}

func PartTwo(rules [][]int, pqs [][]int) int {
	pm := priorityMap(rules)
	_, iqs := sortedQueues(pm, pqs)
	for _, q := range iqs {
		fixQ(pm, q)
	}
	total := totalOfMiddles(iqs)
	return total
}

func fixQ(pm map[int][]int, q []int) {
	l := len(q)
	for i := l - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			cur := q[j]
			nxt := q[j+1]
			before, ex := pm[cur]
			if ex {
				in := contains(nxt, before)
				if in {
					t := q[j]
					q[j] = q[j+1]
					q[j+1] = t
				}
			}
		}
	}
}

func totalOfMiddles(cqs [][]int) int {
	total := 0
	for _, c := range cqs {
		m := len(c) / 2
		total += c[m]
	}
	return total
}

func sortedQueues(pm map[int][]int, pqs [][]int) ([][]int, [][]int) {
	cqs := [][]int{}
	iqs := [][]int{}
	for i := 0; i < len(pqs); i++ {
		correct := true
		for j := 0; j < len(pqs[i]); j++ {
			cur := pqs[i][j]
			before, ex := pm[cur]
			if ex {
				for k := j; k < len(pqs[i]); k++ {
					p := pqs[i][k]
					in := contains(p, before)
					if in {
						correct = false
						break
					}
				}
			}
			if !correct {
				break
			}
		}
		if correct {
			cqs = append(cqs, pqs[i])
		} else {
			iqs = append(iqs, pqs[i])
		}
	}
	return cqs, iqs
}

func contains(n int, s []int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func priorityMap(rules [][]int) map[int][]int {
	pm := make(map[int][]int)
	for _, v := range rules {
		pm[v[1]] = append(pm[v[1]], v[0])
	}
	return pm
}
