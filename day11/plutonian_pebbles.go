package day11

import (
	"container/list"
	"strconv"
)

// Part Two would work for Part One, but Part One solution had memory
// issues when blinks value got higher. Part Two is better solution.

func PartOne(lst *list.List, blinks int) int {
	for range blinks {
		for e := lst.Front(); e != nil; e = e.Next() {
			var l, r string
			if v, ok := e.Value.(string); ok {
				l, r = applyRule(v)
			} else {
				panic("list value not string")
			}
			if r != "-1" {
				e.Value = l
				e = lst.InsertAfter(r, e)
			} else {
				e.Value = l
			}
		}
	}
	return lst.Len()
}

func PartTwo(smap map[string]int, blinks int) int {
	for range blinks {
		temp := make(map[string]int)
		for s, c := range smap {
			ns := applyRuleTwo(s)
			for _, v := range ns {
				temp[v] += c
			}
		}
		smap = temp
	}
	total := 0
	for s := range smap {
		total += smap[s]
	}
	return total
}

func applyRuleTwo(s string) []string {
	if s == "0" {
		return []string{"1"}
	} else if len(s)%2 == 0 {
		m := len(s) / 2
		l := s[:m]
		r := s[m:]
		r = stripZero(r)
		return []string{l, r}

	} else {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		v := strconv.Itoa(i * 2024)
		return []string{v}
	}
}

func applyRule(s string) (string, string) {
	if s == "0" {
		return "1", "-1"
	} else if len(s)%2 == 0 {
		m := len(s) / 2
		l := s[:m]
		r := s[m:]
		r = stripZero(r)
		return l, r

	} else {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		v := strconv.Itoa(i * 2024)
		return v, "-1"
	}
}

func stripZero(s string) string {
	l := len(s)
	trim := 0
	for i := 0; i < l-1; i++ {
		if s[i] == 48 {
			trim++
		} else {
			break
		}
	}
	return s[trim:]
}
