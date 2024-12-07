package day4

import (
	"testing"
)

func TestCeresSearchAoCInput(t *testing.T) {
	file := true
	cs := [][]string{}
	total := CeresSearch(file, cs)
	if total != 2543 {
		t.Errorf("Ceres Search total = %d, expected 3", total)
	}
}

func TestCeresSearch(t *testing.T) {
	file := false
	cs := [][]string{
		{"X", "M", "A", "S"},
		{"M", "M", "S", "A"},
		{"A", "S", "A", "X"},
		{"S", "X", "M", "S"},
	}
	total := CeresSearch(file, cs)
	if total != 3 {
		t.Errorf("Ceres Search total = %d, expected 3", total)
	}
}

func TestXmasAoCInput(t *testing.T) {
	file := true
	cs := [][]string{}
	total := Xmas(file, cs)
	if total != 1930 {
		t.Errorf("Xmas Cross total = %d, expected 1", total)
	}
}

func TestXmas(t *testing.T) {
	file := false
	cs := [][]string{
		{"M", ".", "M"},
		{".", "A", "."},
		{"S", ".", "S"},
	}
	total := Xmas(file, cs)
	if total != 1 {
		t.Errorf("Xmas Cross total = %d, expected 1", total)
	}
}
