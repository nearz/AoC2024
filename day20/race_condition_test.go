package day20

import (
	"fmt"
	"testing"
)

func TestPartOneAocInput(t *testing.T) {
	rcmap := RCFile("rc.input")
	total := PartOne(rcmap)
	fmt.Println(total)
}

func TestPartOne(t *testing.T) {
	rcmap := RCFile("t1.input")
	total := PartOne(rcmap)
	fmt.Println(total)
}
