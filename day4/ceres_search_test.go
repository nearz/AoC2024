package day4

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	cs, err := CeresSearchFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(cs)
	if total != 2543 {
		t.Errorf("Ceres Search total = %d, expected 2543", total)
	}
}

func TestPartOne(t *testing.T) {
	cs := [][]string{
		{"X", "M", "A", "S"},
		{"M", "M", "S", "A"},
		{"A", "S", "A", "X"},
		{"S", "X", "M", "S"},
	}
	total := PartOne(cs)
	if total != 3 {
		t.Errorf("Ceres Search total = %d, expected 3", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	cs, err := CeresSearchFile()
	if err != nil {
		t.Error(err)
	}
	total := PartTwo(cs)
	if total != 1930 {
		t.Errorf("Xmas Cross total = %d, expected 1", total)
	}
}

func TestPartTwo(t *testing.T) {
	cs := [][]string{
		{"M", ".", "M"},
		{".", "A", "."},
		{"S", ".", "S"},
	}
	total := PartTwo(cs)
	if total != 1 {
		t.Errorf("Xmas Cross total = %d, expected 1", total)
	}
}
