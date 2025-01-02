package day18

import (
	"fmt"
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	pslc := RRFile("rr.input")
	total := PartOne(pslc, finalGrid, bLen)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	pslc := RRFile("t1.input")
	total := PartOne(pslc, testGrid, tbLen)
	if total != 22 {
		t.Errorf("Part One total = %d, expected 22", total)
	}
}

func TestPartTwo(t *testing.T) {
	pslc := RRFile("t1.input")
	blockPos := PartTwo(pslc, testGrid)
	fmt.Println(blockPos)
}

func TestPartTwoAoCInput(t *testing.T) {
	pslc := RRFile("rr.input")
	blockPos := PartTwo(pslc, finalGrid)
	fmt.Println(blockPos)
}
