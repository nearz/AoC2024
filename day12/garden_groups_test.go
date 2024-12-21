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
