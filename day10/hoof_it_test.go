package day10

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	tmap, err := HoofItFile()
	if err != nil {
		t.Error(err)
		return
	}
	total := PartOne(tmap)
	if total != 644 {
		t.Errorf("Part One total = %d, expected 644", total)
	}
}

func TestPartOne(t *testing.T) {
	tmap := [][]int{
		{-1, -1, -1, 0, -1, -1, -1},
		{-1, -1, -1, 1, -1, -1, -1},
		{-1, -1, -1, 2, -1, -1, -1},
		{6, 5, 4, 3, 4, 5, 6},
		{7, 6, 5, 4, -1, -1, 7},
		{8, -1, -1, -1, -1, -1, 8},
		{9, -1, -1, -1, -1, -1, 9},
	}
	total := PartOne(tmap)
	if total != 2 {
		t.Errorf("Part One total = %d, expected 2", total)
	}

	tmap2 := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}
	total = PartOne(tmap2)
	if total != 36 {
		t.Errorf("Part One total = %d, expected 36", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	tmap, err := HoofItFile()
	if err != nil {
		t.Error(err)
		return
	}
	total := PartTwo(tmap)
	if total != 1366 {
		t.Errorf("Part One total = %d, expected 1366", total)
	}
}

func TestPartTwo(t *testing.T) {
	tmap2 := [][]int{
		{8, 9, 0, 1, 0, 1, 2, 3},
		{7, 8, 1, 2, 1, 8, 7, 4},
		{8, 7, 4, 3, 0, 9, 6, 5},
		{9, 6, 5, 4, 9, 8, 7, 4},
		{4, 5, 6, 7, 8, 9, 0, 3},
		{3, 2, 0, 1, 9, 0, 1, 2},
		{0, 1, 3, 2, 9, 8, 0, 1},
		{1, 0, 4, 5, 6, 7, 3, 2},
	}
	total := PartTwo(tmap2)
	if total != 81 {
		t.Errorf("Part One total = %d, expected 81", total)
	}
}
