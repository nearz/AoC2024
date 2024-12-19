package day11

import (
	"container/list"
	"testing"
)

func TestPartTwoAoCInput(t *testing.T) {
	smap := make(map[string]int)
	smap["0"] = 1
	smap["7"] = 1
	smap["198844"] = 1
	smap["5687836"] = 1
	smap["58"] = 1
	smap["2478"] = 1
	smap["25475"] = 1
	smap["894"] = 1
	total := PartTwo(smap, 75)
	if total != 257335372288947 {
		t.Errorf("Part Two total = %d, expected 257335372288947", total)
	}
}

func TestPartTwo(t *testing.T) {
	smap := make(map[string]int)
	smap["125"] = 1
	smap["17"] = 1
	total := PartTwo(smap, 25)
	if total != 55312 {
		t.Errorf("Part One total = %d, expected 55312", total)
	}
}

func TestPartOneAoCInput(t *testing.T) {
	l := list.New()
	l.PushBack("0")
	l.PushBack("7")
	l.PushBack("198844")
	l.PushBack("5687836")
	l.PushBack("58")
	l.PushBack("2478")
	l.PushBack("25475")
	l.PushBack("894")
	total := PartOne(l, 25)
	if total != 216996 {
		t.Errorf("Part One total = %d, expected 216996", total)
	}
}

func TestPartOne(t *testing.T) {
	l := list.New()
	l.PushBack("125")
	l.PushBack("17")
	total := PartOne(l, 25)
	if total != 55312 {
		t.Errorf("Part One total = %d, expected 55312", total)
	}
}
