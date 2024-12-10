package day3

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	s, err := MullItOverFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(s)
	if total != 155955228 {
		t.Errorf("Mull It Over = %d, expected 155955228", total)
	}
}

func TestPartOne(t *testing.T) {
	s := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	total := PartOne(s)
	if total != 161 {
		t.Errorf("Test MIO = %d, expected 161", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	s, err := MullItOverFile()
	if err != nil {
		t.Error(err)
	}
	total := PartTwo(s)
	if total != 100189366 {
		t.Errorf("Dos and Donts = %d, expected 100189366", total)
	}
}

func TestPartTwo(t *testing.T) {
	s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	total := PartTwo(s)
	if total != 48 {
		t.Errorf("Dos and Donts = %d, expected 48", total)
	}
}
