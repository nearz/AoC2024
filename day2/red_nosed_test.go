package day2

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	reports, err := RedNosedFile()
	if err != nil {
		t.Error(err)
	}
	safetyCount := PartOne(reports)
	if safetyCount["safe"] != 213 {
		t.Errorf("Safe count = %d, expected 213", safetyCount["safe"])
	}

	if safetyCount["unsafe"] != 787 {
		t.Errorf("Safe count = %d, expected 787", safetyCount["unsafe"])
	}
}

func TestPartOne(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	safetyCount := PartOne(reports)
	if safetyCount["safe"] != 2 {
		t.Errorf("Safe count = %d, expected 2", safetyCount["safe"])
	}
	if safetyCount["unsafe"] != 4 {
		t.Errorf("Unsafe count = %d, expected 4", safetyCount["unsafe"])
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	reports, err := RedNosedFile()
	if err != nil {
		t.Error(err)
	}
	safetyCount := PartTwo(reports)
	if safetyCount["safe"] != 285 {
		t.Errorf("Safe count = %d, expected 285", safetyCount["safe"])
	}
	if safetyCount["unsafe"] != 715 {
		t.Errorf("Safe count = %d, expected 715", safetyCount["unsafe"])
	}
}

func TestPartTwo(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	safetyCount := PartTwo(reports)
	if safetyCount["safe"] != 4 {
		t.Errorf("Safe count = %d, expected 4", safetyCount["safe"])
	}
	if safetyCount["unsafe"] != 2 {
		t.Errorf("Unsafe count = %d, expected 2", safetyCount["unsafe"])
	}
}
