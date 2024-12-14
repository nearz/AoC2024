package day7

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	eqs, err := BridgeRepairFile()
	if err != nil {
		t.Error(err)
	}
	total := PartOne(eqs)
	if total != 6083020304036 {
		t.Errorf("Part One = %d, expected 6083020304036", total)
	}
}

func TestPartOne(t *testing.T) {
	eqs := []Equation{
		{res: 190, operands: []int{10, 19}},
		{res: 3267, operands: []int{81, 40, 27}},
		{res: 83, operands: []int{17, 5}},
		{res: 156, operands: []int{15, 6}},
		{res: 7290, operands: []int{6, 8, 6, 15}},
		{res: 161011, operands: []int{16, 10, 13}},
		{res: 192, operands: []int{17, 8, 14}},
		{res: 21037, operands: []int{9, 7, 18, 13}},
		{res: 292, operands: []int{11, 6, 16, 20}},
	}
	total := PartOne(eqs)
	if total != 3749 {
		t.Errorf("Part One = %d, expected 3749", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	eqs, err := BridgeRepairFile()
	if err != nil {
		t.Error(err)
	}
	total := PartTwo(eqs)
	if total != 59002246504791 {
		t.Errorf("Part One = %d, expected 59002246504791", total)
	}
}

func TestPartTwo(t *testing.T) {
	eqs := []Equation{
		{res: 190, operands: []int{10, 19}},
		{res: 3267, operands: []int{81, 40, 27}},
		{res: 83, operands: []int{17, 5}},
		{res: 156, operands: []int{15, 6}},
		{res: 7290, operands: []int{6, 8, 6, 15}},
		{res: 161011, operands: []int{16, 10, 13}},
		{res: 192, operands: []int{17, 8, 14}},
		{res: 21037, operands: []int{9, 7, 18, 13}},
		{res: 292, operands: []int{11, 6, 16, 20}},
	}
	total := PartTwo(eqs)
	if total != 11387 {
		t.Errorf("Part One = %d, expected 11387", total)
	}
}
