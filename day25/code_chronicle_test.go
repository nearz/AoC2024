package day25

import (
	"fmt"
	"testing"
)

func TestPartOneAocInput(t *testing.T) {
	schematic := CCFile("cc.input")
	total := PartOne(schematic)
	fmt.Println(total)
}

// func TestPartOne(t *testing.T) {
// 	schematic := CCFile("t1.input")
// 	total := PartOne(schematic)
// 	if total != 3 {
// 		t.Errorf("Part One total = %d, expected 3", total)
// 	}
// }
