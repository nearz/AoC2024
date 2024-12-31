package day19

import (
	"strings"
)

func PartOne(twls []string, dsgns []string) int {
	total := 0
	for _, d := range dsgns {
		mem := make(map[string]int)
		r := builds(d, twls, mem)
		if r > 0 {
			total++
		}
	}
	return total
}

func PartTwo(twls []string, dsgns []string) int {
	total := 0
	for _, d := range dsgns {
		mem := make(map[string]int)
		r := builds(d, twls, mem)
		total += r
	}
	return total
}

func builds(dsgn string, twls []string, mem map[string]int) (cnt int) {
	if v, ok := mem[dsgn]; ok {
		return v
	}
	for _, t := range twls {
		if t == dsgn {
			cnt++
		} else if strings.HasPrefix(dsgn, t) {
			cnt += builds(strings.TrimPrefix(dsgn, t), twls, mem)
		}
	}
	mem[dsgn] = cnt
	return
}
