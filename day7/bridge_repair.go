package day7

import "strconv"

type Equation struct {
	res      int
	operands []int
}

func PartOne(eqs []Equation) int {
	total := 0
	ops := []string{"*", "+"}
	for _, e := range eqs {
		total += checkEq(e, ops)
	}
	return total
}

func PartTwo(eqs []Equation) int {
	total := 0
	ops := []string{"*", "+", "|"}
	for _, e := range eqs {
		total += checkEq(e, ops)
	}
	return total
}

func checkEq(e Equation, ops []string) int {
	n := len(e.operands)
	ps := permuts(ops, n-1)
	for _, p := range ps {
		v := calc(p, e.operands)
		if v == e.res {
			return e.res
		}
	}
	return 0
}

func calc(ops string, opds []int) int {
	if len(ops) == 0 {
		return opds[0]
	}
	lops := len(ops) - 1
	lopds := len(opds) - 1
	if ops[lops] == 42 {
		return calc(ops[:lops], opds[:lopds]) * opds[lopds]
	} else if ops[lops] == 124 {
		f := strconv.Itoa(calc(ops[:lops], opds[:lopds]))
		s := strconv.Itoa(opds[lopds])
		r, err := strconv.Atoi(f + s)
		if err != nil {
			panic(err)
		}
		return r
	}
	return calc(ops[:lops], opds[:lopds]) + opds[lopds]
}

func permuts(operators []string, n int) []string {
	if n == 0 {
		return []string{""}
	}

	oneLess := permuts(operators, n-1)
	var allPermuts []string
	for _, ol := range oneLess {
		for _, o := range operators {
			allPermuts = append(allPermuts, ol+o)
		}
	}
	return allPermuts
}
