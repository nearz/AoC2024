package day24

import (
	"sort"
)

type wire struct {
	name string
	val  int
}

type gate struct {
	in1   wire
	in2   wire
	out   wire
	logic string
}

func PartOne(active map[string]int, gates []gate) int {
	zstate := processGates(active, gates)
	btd := binToDec(zstate)
	return btd
}

func processGates(active map[string]int, gates []gate) []wire {
	zstate := []wire{}
	for len(gates) > 0 {
		cg := gates[0]
		inv1, inok1 := active[cg.in1.name]
		inv2, inok2 := active[cg.in2.name]
		if inok1 && inok2 {
			cg.in1.val = inv1
			cg.in2.val = inv2
			switch cg.logic {
			case "AND":
				cg.out.val = cg.in1.val & cg.in2.val
			case "OR":
				cg.out.val = cg.in1.val | cg.in2.val
			case "XOR":
				cg.out.val = cg.in1.val ^ cg.in2.val
			}
			if string(cg.out.name[0]) == "z" {
				zstate = append(zstate, cg.out)
			} else {
				active[cg.out.name] = cg.out.val
			}
			gates = gates[1:]
		} else {
			gates = gates[1:]
			gates = append(gates, cg)
		}
	}
	sort.Slice(zstate, func(i, j int) bool {
		return zstate[i].name < zstate[j].name
	})
	return zstate
}

func binToDec(zstate []wire) int {
	total := 0
	zl := len(zstate)
	for i := 0; i < zl; i++ {
		total += zstate[i].val * intPow(2, i)
	}
	return total
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}
