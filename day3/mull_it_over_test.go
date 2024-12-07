package day3

import (
	"testing"
)

func TestMullItOverAoCInput(t *testing.T) {
	file := true
	s := ""
	total := MullItOver(file, s)
	if total != 155955228 {
		t.Errorf("Mull It Over = %d, expected 155955228", total)
	}
}

func TestMullItOver(t *testing.T) {
	file := false
	s := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	total := MullItOver(file, s)
	if total != 161 {
		t.Errorf("Test MIO = %d, expected 161", total)
	}
}

func TestDosAndDontsAoCInput(t *testing.T) {
	file := true
	s := ""
	total := DosAndDonts(file, s)
	if total != 100189366 {
		t.Errorf("Dos and Donts = %d, expected 100189366", total)
	}
}

func TestDosAndDonts(t *testing.T) {
	file := false
	s := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	total := DosAndDonts(file, s)
	if total != 48 {
		t.Errorf("Dos and Donts = %d, expected 48", total)
	}
}
