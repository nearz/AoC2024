package day9

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	dm, err := DFFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(dm)
	if total != 6382875730645 {
		t.Errorf("Part One total = %d, expected 6382875730645", total)
	}
}

func TestPartOne(t *testing.T) {
	dm := "2333133121414131402"
	total := PartOne(dm)
	if total != 1928 {
		t.Errorf("Part One total = %d, expected 1928", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	dm, err := DFFile()
	if err != nil {
		t.Error(err)
	}
	total := PartTwo(dm)
	if total != 6420913943576 {
		t.Errorf("Part One total = %d, expected 6420913943576", total)
	}
}

func TestPartTwo(t *testing.T) {
	dm := "2333133121414131402"
	total := PartTwo(dm)
	if total != 2858 {
		t.Errorf("Part One total = %d, expected 2858", total)
	}
}
