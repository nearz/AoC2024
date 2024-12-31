package day19

import (
	"fmt"
	"testing"
)

func TestPartTwoAoCInput(t *testing.T) {
	twls, dsgns := LLFile("ll.input")
	total := PartTwo(twls, dsgns)
	fmt.Println(total)
}

func TestPartOneAoCInput(t *testing.T) {
	twls, dsgns := LLFile("ll.input")
	total := PartOne(twls, dsgns)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	twls, dsgns := LLFile("t1.input")
	total := PartOne(twls, dsgns)
	if total != 6 {
		t.Errorf("Part One total = %d, expected 6", total)
	}
}
