package day23

import (
	"fmt"
	"testing"
)

func TestPartOneAocInput(t *testing.T) {
	cprs := LPFile("lp.input")
	total := PartOne(cprs)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	cprs := LPFile("t1.input")
	total := PartOne(cprs)
	if total != 7 {
		t.Errorf("Part One total = %d, expected 7", total)
	}
}

// TBC
// func TestPartTwo(t *testing.T) {
// 	cprs := LPFile("t1.input")
// 	psswrd := PartTwo(cprs)
// 	fmt.Println(psswrd)
// }
