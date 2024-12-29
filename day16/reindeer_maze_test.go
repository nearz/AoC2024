package day16

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	rmap := RMFile("t1.input")
	total := PartOne(rmap)
	if total != 7036 {
		t.Errorf("Part One total = %d, expected 7036", total)
	}
}
