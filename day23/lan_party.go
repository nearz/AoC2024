package day23

import (
	"sort"
	"strings"
)

type cpair struct {
	c1 string
	c2 string
}

type (
	cset    map[string]struct{}
	adjList map[string][]string
)

func PartOne(cprs []cpair) int {
	al := adjacencyList(cprs)
	tss := findThreeSet(al)
	total := 0
	for _, ts := range tss {
		for _, c := range ts {
			if string(c[0]) == "t" {
				total++
				break
			}
		}
	}
	return total
}

// TBC -- Biggest Clique --
// func PartTwo(cprs []cpair) string {
// 	al := adjacencyList(cprs)
// 	fmt.Println(al)
// 	return ""
// }

func findThreeSet(al adjList) [][]string {
	tss := [][]string{}
	vis := make(map[string]struct{})
	for c, cps := range al {
		cl := len(cps)
		for i := 0; i < cl; i++ {
			for j := 0; j < cl; j++ {
				if j == i {
					continue
				}
				cnts := contains(cps[j], al[cps[i]])
				if cnts {
					ts := []string{c, cps[i], cps[j]}
					sort.Strings(ts)
					k := strings.Join(ts, "")
					if _, vok := vis[k]; !vok {
						tss = append(tss, ts)
						vis[k] = struct{}{}
					}
				}
			}
		}
	}
	return tss
}

func contains(v string, ss []string) bool {
	for _, val := range ss {
		if v == val {
			return true
		}
	}
	return false
}

func adjacencyList(cprs []cpair) adjList {
	al := make(adjList)
	for _, c := range cprs {
		al[c.c1] = append(al[c.c1], c.c2)
		al[c.c2] = append(al[c.c2], c.c1)
	}
	return al
}
