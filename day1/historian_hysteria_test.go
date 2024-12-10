package day1

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	left, right, err := HysteriaFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(left, right)
	if total != 1666427 {
		t.Errorf("Hysteria Total = %d, expected 1666427", total)
	}
}

func TestPartOne(t *testing.T) {
	left := []int{5, 8, 1, 2}
	right := []int{4, 9, 2, 7}
	total := PartOne(left, right)
	if total != 6 {
		t.Errorf("Hysteria Total = %d, expected 6", total)
	}
}

func TestPartTwo(t *testing.T) {
	left := []int{3, 3, 3}
	right := []int{3, 3, 3}
	score := PartTwo(left, right)
	if score != 27 {
		t.Errorf("Similarity Score = %d, expected 27", score)
	}

	leftTwo := []int{2, 2, 3, 4, 5}
	rightTwo := []int{2, 3, 3, 4, 4}
	score = PartTwo(leftTwo, rightTwo)
	if score != 18 {
		t.Errorf("Similarity Score = %d, expected 18", score)
	}
}

func TestPartTwoAoCInputs(t *testing.T) {
	left, right, err := HysteriaFile()
	if err != nil {
		t.Error(err)
	}
	score := PartTwo(left, right)
	if score != 24316233 {
		t.Errorf("Similarity Score = %d, expected 24316233", score)
	}
}
