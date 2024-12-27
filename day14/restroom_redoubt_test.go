package day14

import (
	"testing"
)

func TestPartOneAoCInput(t *testing.T) {
	bots := RRInput("rr.input")
	total := PartOne(bots, finalGrid)
	if total != 219150360 {
		t.Errorf("Part One total = %d, expected 219150360", total)
	}
}

func TestPartOne(t *testing.T) {
	bots := RRInput("t1.input")
	total := PartOne(bots, testGrid)
	if total != 12 {
		t.Errorf("Part One total = %d, expected 12", total)
	}
}

func TestPartTwo(t *testing.T) {
	bots := RRInput("rr.input")
	PartTwo(bots, finalGrid)
}

// Print the christmas tree
func TestPrintTree(t *testing.T) {
	bots := RRInput("rr.input")
	for i := range bots {
		moveBotSeconds(&bots[i], 8053, finalGrid)
	}
	printBots(bots, finalGrid)
}
