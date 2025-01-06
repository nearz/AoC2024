package day24

import (
	"fmt"
	"testing"
)

func TestPartOneAocInput(t *testing.T) {
	init, gates := CWFile("cw.input")
	total := PartOne(init, gates)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	init, gates := CWFile("t1.input")
	total := PartOne(init, gates)
	if total != 2024 {
		t.Errorf("Part One total = %d, expected 2024", total)
	}
}
