package day12

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	garden, err := GardenGroupsFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(garden)
	if total != 1461752 {
		t.Errorf("Part One total = %d, expected 1461752", total)
	}
}

func TestPartOne(t *testing.T) {
	garden := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	total := PartOne(garden)
	if total != 140 {
		t.Errorf("Part One total = %d, expected 140", total)
	}

	garden2 := [][]string{
		{"R", "R", "R", "R", "I", "I", "C", "C", "F", "F"},
		{"R", "R", "R", "R", "I", "I", "C", "C", "C", "F"},
		{"V", "V", "R", "R", "R", "C", "C", "F", "F", "F"},
		{"V", "V", "R", "C", "C", "C", "J", "F", "F", "F"},
		{"V", "V", "V", "V", "C", "J", "J", "C", "F", "E"},
		{"V", "V", "I", "V", "C", "C", "J", "J", "E", "E"},
		{"V", "V", "I", "I", "I", "C", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "I", "I", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "S", "I", "J", "E", "E", "E"},
		{"M", "M", "M", "I", "S", "S", "J", "E", "E", "E"},
	}
	total = PartOne(garden2)
	if total != 1930 {
		t.Errorf("Part One total = %d, expected 1930", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	garden, err := GardenGroupsFile()
	if err != nil {
		t.Error(err)
	}
	total := PartTwo(garden)
	if total != 904114 {
		t.Errorf("Part One total = %d, expected 904114", total)
	}
}

func TestPartTwo(t *testing.T) {
	garden := [][]string{
		{"A", "A", "A", "A"},
		{"B", "B", "C", "D"},
		{"B", "B", "C", "C"},
		{"E", "E", "E", "C"},
	}
	total := PartTwo(garden)
	if total != 80 {
		t.Errorf("Part Two total = %d, expected 80", total)
	}

	g2 := [][]string{
		{"E", "E", "E", "E", "E"},
		{"E", "X", "X", "X", "X"},
		{"E", "E", "E", "E", "E"},
		{"E", "X", "X", "X", "X"},
		{"E", "E", "E", "E", "E"},
	}
	t2 := PartTwo(g2)
	if t2 != 236 {
		t.Errorf("Part Two total = %d, expected 236", t2)
	}

	g3 := [][]string{
		{"A", "A", "A", "A", "A", "A"},
		{"A", "A", "A", "B", "B", "A"},
		{"A", "A", "A", "B", "B", "A"},
		{"A", "B", "B", "A", "A", "A"},
		{"A", "B", "B", "A", "A", "A"},
		{"A", "A", "A", "A", "A", "A"},
	}
	t3 := PartTwo(g3)
	if t3 != 368 {
		t.Errorf("Part Two total = %d, expected 368", t3)
	}

	g4 := [][]string{
		{"O", "O", "O", "O", "O"},
		{"O", "X", "O", "X", "O"},
		{"O", "O", "O", "O", "O"},
		{"O", "X", "O", "X", "O"},
		{"O", "O", "O", "O", "O"},
	}
	t4 := PartTwo(g4)
	if t4 != 436 {
		t.Errorf("Part Two total = %d, expected 436", t4)
	}

	g5 := [][]string{
		{"R", "R", "R", "R", "I", "I", "C", "C", "F", "F"},
		{"R", "R", "R", "R", "I", "I", "C", "C", "C", "F"},
		{"V", "V", "R", "R", "R", "C", "C", "F", "F", "F"},
		{"V", "V", "R", "C", "C", "C", "J", "F", "F", "F"},
		{"V", "V", "V", "V", "C", "J", "J", "C", "F", "E"},
		{"V", "V", "I", "V", "C", "C", "J", "J", "E", "E"},
		{"V", "V", "I", "I", "I", "C", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "I", "I", "J", "J", "E", "E"},
		{"M", "I", "I", "I", "S", "I", "J", "E", "E", "E"},
		{"M", "M", "M", "I", "S", "S", "J", "E", "E", "E"},
	}
	t5 := PartTwo(g5)
	if t5 != 1206 {
		t.Errorf("Part One total = %d, expected 1206", t5)
	}
}
