package day2

import (
	"testing"
)

func TestRedNosedAoCInput(t *testing.T) {
	file := true
	reports := [][]int{}
	safetyCount := RedNosed(file, reports)
	if safetyCount["safe"] != 213 {
		t.Errorf("Safe count = %d, expected 213", safetyCount["safe"])
	}

	if safetyCount["unsafe"] != 787 {
		t.Errorf("Safe count = %d, expected 787", safetyCount["unsafe"])
	}
}

func TestRedNosed(t *testing.T) {
	file := false
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	safetyCount := RedNosed(file, reports)
	if safetyCount["safe"] != 2 {
		t.Errorf("Safe count = %d, expected 2", safetyCount["safe"])
	}
	if safetyCount["unsafe"] != 4 {
		t.Errorf("Unsafe count = %d, expected 4", safetyCount["unsafe"])
	}
}

func TestRedNosedDampenedAoCInput(t *testing.T) {
	file := true
	reports := [][]int{}
	safetyCount := RedNosedDampened(file, reports)
	if safetyCount["safe"] != 285 {
		t.Errorf("Safe count = %d, expected 285", safetyCount["safe"])
	}
	if safetyCount["unsafe"] != 715 {
		t.Errorf("Safe count = %d, expected 715", safetyCount["unsafe"])
	}
}

// func TestTemp(t *testing.T) {
// 	file := false
// 	r := [][]int{
// 		{69, 71, 69, 66, 63},
// 		{15, 13, 15, 16, 17, 20, 21, 24},
// 	}
// 	c := RedNosedDampened(file, r)
// 	fmt.Println(c)
// }

//
// func TestUnsafe(t *testing.T) {
// 	r := []int{69, 71, 69, 66, 63}
// 	u := unsafeLevel(r)
// 	fmt.Println(u)
// }

func TestRedNosedDampened(t *testing.T) {
	file := false
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	safetyCount := RedNosedDampened(file, reports)
	if safetyCount["safe"] != 4 {
		t.Errorf("Safe count = %d, expected 4", safetyCount["safe"])
	}
	if safetyCount["unsafe"] != 2 {
		t.Errorf("Unsafe count = %d, expected 2", safetyCount["unsafe"])
	}
}
