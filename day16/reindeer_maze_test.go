package day16

import (
	"fmt"
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	rmap := RMFile("rm.input")
	total := PartOne(rmap)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	rmap := RMFile("t1.input")
	total := PartOne(rmap)
	if total != 7036 {
		t.Errorf("Part One total = %d, expected 7036", total)
	}

	rmap = RMFile("t2.input")
	total = PartOne(rmap)
	if total != 11048 {
		t.Errorf("Part One total = %d, expected 11048", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	rmap := RMFile("rm.input")
	total := PartTwo(rmap)
	fmt.Println(total)
}

func TestPartTwo(t *testing.T) {
	rmap := RMFile("t1.input")
	total := PartTwo(rmap)
	if total != 45 {
		t.Errorf("Part Two total = %d, expected 45", total)
	}

	rmap = RMFile("t2.input")
	total = PartTwo(rmap)
	if total != 64 {
		t.Errorf("Part Two total = %d, expected 64", total)
	}
}
