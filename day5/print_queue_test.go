package day5

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	rules := Rules()
	pqs := ActivePrintQueues()
	total := PartOne(rules, pqs)
	if total != 5732 {
		t.Errorf("Printer Q Aoc Input total = %d, expected 5732", total)
	}
}

func TestPartOne(t *testing.T) {
	rules := [][]int{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}
	pqs := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	total := PartOne(rules, pqs)
	if total != 143 {
		t.Errorf("Print Queues total = %d, expected 143", total)
	}
}

func TestPartTwoAoCInput(t *testing.T) {
	rules := Rules()
	pqs := ActivePrintQueues()
	total := PartTwo(rules, pqs)
	if total != 4716 {
		t.Errorf("Printer Q Aoc Input total = %d, expected 4716", total)
	}
}

func TestPartTwo(t *testing.T) {
	rules := [][]int{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}
	pqs := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	total := PartTwo(rules, pqs)
	if total != 123 {
		t.Errorf("Print Queues total = %d, expected 123", total)
	}
}
