package day8

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	amap, err := RCFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(amap)
	if total != 249 {
		t.Errorf("Part One total = %d, expected 249", total)
	}
}

func TestPartOne(t *testing.T) {
	amap := [][]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", "a", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "a", "."},
		{".", ".", ".", ".", ".", "a", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	total := PartOne(amap)
	if total != 4 {
		t.Errorf("Part One total = %d, expected 4", total)
	}

	amap2 := [][]string{
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "0", ".", ".", "."},
		{".", ".", ".", ".", ".", "0", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "0", ".", ".", ".", "."},
		{".", ".", ".", ".", "0", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", "A", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "A", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "A", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
	total = PartOne(amap2)
	if total != 14 {
		t.Errorf("Part One total = %d, expected 14", total)
	}
}
