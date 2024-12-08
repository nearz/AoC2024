package day5

import "fmt"

func PrintQueue(file bool, paramRules [][]int, paramPqs [][]int) int {
	var rules [][]int
	var pqs [][]int
	if file {
		rules = Rules()
		pqs = ActivePrintQueues()
	} else {
		rules = paramRules
		pqs = paramPqs
	}

	pm := priorityMap(rules)
	cqs, _ := sortedQueues(pm, pqs)
	total := totalOfMiddles(cqs)
	return total
}

func CorrectTheQueues(file bool, paramRules [][]int, paramPqs [][]int) int {
	var rules [][]int
	var pqs [][]int
	if file {
		rules = Rules()
		pqs = ActivePrintQueues()
	} else {
		rules = paramRules
		pqs = paramPqs
	}

	pm := priorityMap(rules)
	_, iqs := sortedQueues(pm, pqs)
	fmt.Println(iqs)
	// total := totalOfMiddles(cqs)
	return -1
}

func totalOfMiddles(cqs [][]int) int {
	total := 0
	for _, c := range cqs {
		total += middleOf(c)
	}
	return total
}

func middleOf(s []int) int {
	m := len(s) / 2
	return s[m]
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
