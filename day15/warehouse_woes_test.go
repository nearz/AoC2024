package day15

import (
	"fmt"
	"testing"
)

func TestPartOneAocInput(t *testing.T) {
	wmap, dirStr := WWFile("ww.input")
	total := PartOne(wmap, dirStr)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	wmap, dirStr := WWFile("t1.input")
	total := PartOne(wmap, dirStr)
	if total != 10092 {
		t.Errorf("Part One total = %d, expected 10092", total)
	}
}

// func TestPartTwo(t *testing.T) {
// 	wmap, dirStr := WWFile("t4.input")
// 	PartTwo(wmap, dirStr)
// }
