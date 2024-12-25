package day13

import (
	"fmt"
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	games := CCFile("cc.input")
	total := PartOne(games)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	games := CCFile("p1t1.input")
	total := PartOne(games)
	if total != 480 {
		t.Errorf("Part One total = %d, expected 480", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	games := CCFile("cc.input")
	total := PartTwo(games)
	fmt.Println(total)
}

func TestPartTwo(t *testing.T) {
	games := CCFile("p1t1.input")
	total := PartTwo(games)
	fmt.Println(total)
}
