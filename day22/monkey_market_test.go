package day22

import (
	"fmt"
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	snums := MMFile("mm.input")
	total := PartOne(snums)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	snums := MMFile("t1.input")
	total := PartOne(snums)
	if total != 37327623 {
		t.Errorf("Part One total = %d, expected 37327623", total)
	}
}
